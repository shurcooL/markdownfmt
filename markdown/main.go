// Package markdown provides a Markdown renderer.
package markdown

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/shurcooL/go/indentwriter"
	"gopkg.in/russross/blackfriday.v2"
)

var _ blackfriday.Renderer = (*markdownRenderer)(nil)

type markdownRenderer struct {
	normalTextMarker   map[*bytes.Buffer]int
	orderedListCounter map[int]int
	paragraph          map[int]bool // Used to keep track of whether a given list item uses a paragraph for large spacing.
	listDepth          int
	lastNormalText     string

	// TODO: Clean these up.
	headers      []string
	columnAligns []blackfriday.CellAlignFlags
	columnWidths []int
	cells        []string

	opt Options

	// stringWidth is used internally to calculate visual width of a string.
	stringWidth func(s string) (width int)
}

func (mr *markdownRenderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	switch node.Type {
	case blackfriday.BlockQuote:
		mr.BlockQuote(w, node)
		return blackfriday.SkipChildren
	case blackfriday.List:
		mr.List(w, node, entering)
	case blackfriday.Item:
		mr.ListItem(w, node)
		return blackfriday.SkipChildren
	case blackfriday.Paragraph:
		mr.Paragraph(w, entering)
	case blackfriday.Heading:
		return mr.Header(w, node, entering)
	case blackfriday.HorizontalRule:
		mr.HRule(w)
	case blackfriday.Emph:
		io.WriteString(w, "*")
	case blackfriday.Strong:
		mr.Strong(w, entering)
	case blackfriday.Del:
		io.WriteString(w, "~~")
	case blackfriday.Link:
		mr.Link(w, node)
		return blackfriday.SkipChildren
	case blackfriday.Image:
		mr.Image(w, node.LinkData.Destination, node.LinkData.Title, entering)
	case blackfriday.Text:
		mr.NormalText(w, node.Literal)
	case blackfriday.HTMLBlock:
		mr.BlockHtml(w, node.Literal)
	case blackfriday.CodeBlock:
		mr.BlockCode(w, node)
	case blackfriday.Softbreak:
	case blackfriday.Hardbreak:
		io.WriteString(w, "  \n")
	case blackfriday.Code:
		io.WriteString(w, "`")
		w.Write(node.Literal)
		io.WriteString(w, "`")
	case blackfriday.HTMLSpan:
		w.Write(node.Literal)
	case blackfriday.Table:
		if !entering {
			mr.Table(w)
		}
	case blackfriday.TableCell:
		var buf bytes.Buffer
		for n := node.FirstChild; n != nil; n = n.Next {
			n.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
				return mr.RenderNode(&buf, node, entering)
			})
		}

		if node.TableCellData.IsHeader {
			mr.TableHeaderCell(buf.Bytes(), node.TableCellData.Align)
		} else {
			mr.TableCell(buf.Bytes())
		}
		return blackfriday.SkipChildren
	}
	return blackfriday.GoToNext
}

func (_ *markdownRenderer) RenderHeader(w io.Writer, ast *blackfriday.Node) {}
func (_ *markdownRenderer) RenderFooter(w io.Writer, ast *blackfriday.Node) {}

func formatCode(lang string, text []byte) (formattedCode []byte, ok bool) {
	switch lang {
	case "Go", "go":
		gofmt, err := format.Source(text)
		if err != nil {
			return nil, false
		}
		return gofmt, true
	default:
		return nil, false
	}
}

// Block-level callbacks.
func (mr *markdownRenderer) BlockCode(w io.Writer, node *blackfriday.Node) {
	doubleSpace(w)

	lang := string(node.CodeBlockData.Info)

	// Parse out the language name.
	count := 0
	for _, elt := range strings.Fields(lang) {
		if elt[0] == '.' {
			elt = elt[1:]
		}
		if len(elt) == 0 {
			continue
		}
		io.WriteString(w, "```")
		io.WriteString(w, elt)
		count++
		break
	}

	if count == 0 {
		io.WriteString(w, "```")
	}
	io.WriteString(w, "\n")

	if formattedCode, ok := formatCode(lang, node.Literal); ok {
		w.Write(formattedCode)
	} else {
		w.Write(node.Literal)
	}

	io.WriteString(w, "```\n")
}
func (mr *markdownRenderer) BlockQuote(w io.Writer, node *blackfriday.Node) {
	doubleSpace(w)

	var buf bytes.Buffer
	for n := node.FirstChild; n != nil; n = n.Next {
		n.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
			return mr.RenderNode(&buf, node, entering)
		})
	}

	scanner := bufio.NewScanner(&buf)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			io.WriteString(w, ">\n")
			continue
		}
		io.WriteString(w, "> ")
		w.Write(scanner.Bytes())
		io.WriteString(w, "\n")
	}
}
func (_ *markdownRenderer) BlockHtml(w io.Writer, text []byte) {
	doubleSpace(w)
	w.Write(text)
	w.Write([]byte{'\n'})
}
func (mr *markdownRenderer) Header(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if entering {
		doubleSpace(w)
	}

	level := node.HeadingData.Level
	if level >= 3 {
		if entering {
			fmt.Fprint(w, strings.Repeat("#", level), " ")
		} else {
			w.Write([]byte{'\n'})
		}
		return blackfriday.GoToNext
	} else {
		// Write the header to the output using a buffer so we can track how much we write.
		out := withBuffer(w)
		defer out.Flush()

		marker := out.Len()
		for n := node.FirstChild; n != nil; n = n.Next {
			n.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
				return mr.RenderNode(out, node, entering)
			})
		}

		// Track the number of characters written.
		len := mr.stringWidth(string(out.Bytes()[marker:]))
		switch level {
		case 1:
			fmt.Fprint(out, "\n", strings.Repeat("=", len))
		case 2:
			fmt.Fprint(out, "\n", strings.Repeat("-", len))
		}
		io.WriteString(w, "\n")
		return blackfriday.SkipChildren
	}
}
func (_ *markdownRenderer) HRule(w io.Writer) {
	doubleSpace(w)
	io.WriteString(w, "---\n")
}

func (mr *markdownRenderer) List(w io.Writer, node *blackfriday.Node, entering bool) {
	if entering {
		// If we are inside of an item, avoid adding an additional newline.
		if node.Parent == nil || node.Parent.Type != blackfriday.Item {
			doubleSpace(w)
		} else if !node.ListData.Tight {
			// If this list is not tight data, add the newline anyway.
			doubleSpace(w)
		}
		mr.listDepth++
		if node.ListFlags&blackfriday.ListTypeOrdered != 0 {
			mr.orderedListCounter[mr.listDepth] = 1
		}
		mr.paragraph[mr.listDepth] = !node.ListData.Tight
	} else {
		delete(mr.paragraph, mr.listDepth)
		mr.listDepth--
	}
}

func (mr *markdownRenderer) ListItem(w io.Writer, node *blackfriday.Node) {
	if mr.paragraph[mr.listDepth] {
		if node.Prev != nil {
			io.WriteString(w, "\n")
		}
	}

	if node.ListFlags&blackfriday.ListTypeOrdered != 0 {
		fmt.Fprintf(w, "%d.", mr.orderedListCounter[mr.listDepth])
		mr.orderedListCounter[mr.listDepth]++
	} else {
		io.WriteString(w, "-")
	}

	var buf bytes.Buffer
	for n := node.FirstChild; n != nil; n = n.Next {
		n.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
			return mr.RenderNode(&buf, node, entering)
		})
	}
	indentwriter.New(w, 1).Write(buf.Bytes())
}
func (mr *markdownRenderer) Paragraph(w io.Writer, entering bool) {
	if entering {
		doubleSpace(w)
	} else {
		io.WriteString(w, "\n")
	}
}

func (mr *markdownRenderer) Table(w io.Writer) {
	doubleSpace(w)
	for column, cell := range mr.headers {
		io.WriteString(w, "| ")
		io.WriteString(w, cell)
		for i := mr.stringWidth(cell); i < mr.columnWidths[column]; i++ {
			w.Write([]byte{' '})
		}
		w.Write([]byte{' '})
	}
	io.WriteString(w, "|\n")
	for column, width := range mr.columnWidths {
		w.Write([]byte{'|'})
		if mr.columnAligns[column]&blackfriday.TableAlignmentLeft != 0 {
			w.Write([]byte{':'})
		} else {
			w.Write([]byte{'-'})
		}
		w.Write(bytes.Repeat([]byte{'-'}, width))
		if mr.columnAligns[column]&blackfriday.TableAlignmentRight != 0 {
			w.Write([]byte{':'})
		} else {
			w.Write([]byte{'-'})
		}
	}
	io.WriteString(w, "|\n")
	for i := 0; i < len(mr.cells); {
		for column := range mr.headers {
			cell := []byte(mr.cells[i])
			i++
			io.WriteString(w, "| ")
			switch mr.columnAligns[column] {
			default:
				fallthrough
			case blackfriday.TableAlignmentLeft:
				w.Write(cell)
				for i := mr.stringWidth(string(cell)); i < mr.columnWidths[column]; i++ {
					w.Write([]byte{' '})
				}
			case blackfriday.TableAlignmentCenter:
				spaces := mr.columnWidths[column] - mr.stringWidth(string(cell))
				for i := 0; i < spaces/2; i++ {
					w.Write([]byte{' '})
				}
				w.Write(cell)
				for i := 0; i < spaces-(spaces/2); i++ {
					w.Write([]byte{' '})
				}
			case blackfriday.TableAlignmentRight:
				for i := mr.stringWidth(string(cell)); i < mr.columnWidths[column]; i++ {
					w.Write([]byte{' '})
				}
				w.Write(cell)
			}
			w.Write([]byte{' '})
		}
		io.WriteString(w, "|\n")
	}

	mr.headers = nil
	mr.columnAligns = nil
	mr.columnWidths = nil
	mr.cells = nil
}
func (mr *markdownRenderer) TableHeaderCell(text []byte, align blackfriday.CellAlignFlags) {
	mr.columnAligns = append(mr.columnAligns, align)
	columnWidth := mr.stringWidth(string(text))
	mr.columnWidths = append(mr.columnWidths, columnWidth)
	mr.headers = append(mr.headers, string(text))
}
func (mr *markdownRenderer) TableCell(text []byte) {
	columnWidth := mr.stringWidth(string(text))
	column := len(mr.cells) % len(mr.headers)
	if columnWidth > mr.columnWidths[column] {
		mr.columnWidths[column] = columnWidth
	}
	mr.cells = append(mr.cells, string(text))
}

// Span-level callbacks.
func (mr *markdownRenderer) Strong(w io.Writer, entering bool) {
	if entering && mr.opt.Terminal {
		io.WriteString(w, "\x1b[1m") // Bold.
	}
	io.WriteString(w, "**")
	if !entering && mr.opt.Terminal {
		io.WriteString(w, "\x1b[0m") // Reset.
	}
}
func (_ *markdownRenderer) Image(w io.Writer, link []byte, title []byte, entering bool) {
	if entering {
		io.WriteString(w, "![")
	} else {
		io.WriteString(w, "](")
		w.Write(escape(link))
		if len(title) != 0 {
			io.WriteString(w, ` "`)
			w.Write(title)
			io.WriteString(w, `"`)
		}
		io.WriteString(w, ")")
	}
}
func (mr *markdownRenderer) Link(w io.Writer, node *blackfriday.Node) {
	var buf bytes.Buffer
	for n := node.FirstChild; n != nil; n = n.Next {
		n.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
			return mr.RenderNode(&buf, node, entering)
		})
	}

	// There is no title and the destination is the same as the contents.
	// This can be represented as an auto-link.
	if len(node.LinkData.Title) == 0 && bytes.Equal(node.LinkData.Destination, buf.Bytes()) {
		w.Write(escape(node.LinkData.Destination))
		return
	}

	io.WriteString(w, "[")
	w.Write(buf.Bytes())
	io.WriteString(w, "](")
	w.Write(escape(node.LinkData.Destination))
	if len(node.LinkData.Title) != 0 {
		io.WriteString(w, ` "`)
		w.Write(node.LinkData.Title)
		io.WriteString(w, `"`)
	}
	io.WriteString(w, ")")
}
func (_ *markdownRenderer) RawHtmlTag(out *bytes.Buffer, tag []byte) {
	out.Write(tag)
}
func (_ *markdownRenderer) TripleEmphasis(out *bytes.Buffer, text []byte) {
	out.WriteString("***")
	out.Write(text)
	out.WriteString("***")
}
func (_ *markdownRenderer) StrikeThrough(out *bytes.Buffer, text []byte) {
	out.WriteString("~~")
	out.Write(text)
	out.WriteString("~~")
}
func (_ *markdownRenderer) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	out.WriteString("<FootnoteRef: Not implemented.>") // TODO
}

// escape replaces instances of backslash with escaped backslash in text.
func escape(text []byte) []byte {
	return bytes.Replace(text, []byte(`\`), []byte(`\\`), -1)
}

func isNumber(data []byte) bool {
	for _, b := range data {
		if b < '0' || b > '9' {
			return false
		}
	}
	return true
}

func needsEscaping(text []byte, lastNormalText string) bool {
	switch string(text) {
	case `\`,
		"`",
		"*",
		"_",
		"{", "}",
		"[", "]",
		"(", ")",
		"#",
		"+",
		"-":
		return true
	case "!":
		return false
	case ".":
		// Return true if number, because a period after a number must be escaped to not get parsed as an ordered list.
		return isNumber([]byte(lastNormalText))
	case "<", ">":
		return true
	default:
		return false
	}
}

// Low-level callbacks.
func (mr *markdownRenderer) NormalText(w io.Writer, text []byte) {
	normalText := string(text)
	if needsEscaping(text, mr.lastNormalText) {
		text = append([]byte("\\"), text...)
	}
	mr.lastNormalText = normalText
	if mr.listDepth > 0 && string(text) == "\n" { // TODO: See if this can be cleaned up... It's needed for lists.
		return
	}
	cleanString := cleanWithoutTrim(string(text))
	if cleanString == "" {
		return
	}
	if mr.skipSpaceIfNeededNormalText(w, cleanString) { // Skip first space if last character is already a space (i.e., no need for a 2nd space in a row).
		cleanString = cleanString[1:]
	}
	io.WriteString(w, cleanString)
}

func (mr *markdownRenderer) skipSpaceIfNeededNormalText(w io.Writer, cleanString string) bool {
	if cleanString[0] != ' ' {
		return false
	}
	out, ok := w.(interface {
		Bytes() []byte
	})
	if !ok {
		return false
	}
	data := out.Bytes()
	return len(data) > 0 && data[len(data)-1] == ' '
}

// cleanWithoutTrim is like clean, but doesn't trim blanks.
func cleanWithoutTrim(s string) string {
	var b []byte
	var p byte
	for i := 0; i < len(s); i++ {
		q := s[i]
		if q == '\n' || q == '\r' || q == '\t' {
			q = ' '
		}
		if q != ' ' || p != ' ' {
			b = append(b, q)
			p = q
		}
	}
	return string(b)
}

func doubleSpace(w io.Writer) {
	if out, ok := w.(interface {
		Len() int
	}); ok && out.Len() > 0 {
		w.Write([]byte{'\n'})
	}
}

// terminalStringWidth returns width of s, taking into account possible ANSI escape codes
// (which don't count towards string width).
func terminalStringWidth(s string) (width int) {
	width = runewidth.StringWidth(s)
	width -= strings.Count(s, "\x1b[1m") * len("[1m") // HACK, TODO: Find a better way of doing this.
	width -= strings.Count(s, "\x1b[0m") * len("[0m") // HACK, TODO: Find a better way of doing this.
	return width
}

// NewRenderer returns a Markdown renderer.
// If opt is nil the defaults are used.
func NewRenderer(opt *Options) blackfriday.Renderer {
	mr := &markdownRenderer{
		normalTextMarker:   make(map[*bytes.Buffer]int),
		orderedListCounter: make(map[int]int),
		paragraph:          make(map[int]bool),

		stringWidth: runewidth.StringWidth,
	}
	if opt != nil {
		mr.opt = *opt
	}
	if mr.opt.Terminal {
		mr.stringWidth = terminalStringWidth
	}
	return mr
}

// Options specifies options for formatting.
type Options struct {
	// Terminal specifies if ANSI escape codes are emitted for styling.
	Terminal bool
}

// Process formats Markdown.
// If opt is nil the defaults are used.
// Error can only occur when reading input from filename rather than src.
func Process(filename string, src []byte, opt *Options) ([]byte, error) {
	// Get source.
	text, err := readSource(filename, src)
	if err != nil {
		return nil, err
	}

	// extensions for GitHub Flavored Markdown-like parsing.
	const extensions = blackfriday.NoIntraEmphasis |
		blackfriday.Tables |
		blackfriday.FencedCode |
		blackfriday.Autolink |
		blackfriday.Strikethrough |
		blackfriday.SpaceHeadings |
		blackfriday.NoEmptyLineBeforeBlock

	output := blackfriday.Run(text, blackfriday.WithRenderer(NewRenderer(opt)), blackfriday.WithExtensions(extensions))
	return output, nil
}

// If src != nil, readSource returns src.
// If src == nil, readSource returns the result of reading the file specified by filename.
func readSource(filename string, src []byte) ([]byte, error) {
	if src != nil {
		return src, nil
	}
	return ioutil.ReadFile(filename)
}

// buffer ensures that an io.Writer output is buffered either by writing to an
// existing buffer or writing a temporary buffer that will be flushed to the
// underlying io.Writer.
type buffer struct {
	*bytes.Buffer
	w io.Writer
}

// withBuffer ensures that the io.Writer passed in is a buffered output.
// If the io.Writer passed in is a Buffer or bytes.Buffer, it is used directly
// which prevents any memory allocation. If the io.Writer is not a native buffer,
// then the io.Writer is saved and a temporary buffer is created. When Flush is
// called, any data written to the buffer is written to the underlying io.Writer.
//
// This is useful if you want to ensure you are currently working with a buffer
// and allows each individual method to determine if they are appending to an
// existing buffer or need to create their own to perform their own work.
func withBuffer(w io.Writer) buffer {
	if buf, ok := w.(buffer); ok {
		// Copy the bytes.Buffer, but do not copy the writer. This ensures we do
		// not accidentally flush multiple times to the same underlying writer.
		return buffer{Buffer: buf.Buffer}
	} else if buf, ok := w.(*bytes.Buffer); ok {
		return buffer{Buffer: buf}
	}

	// Save the io.Writer and create a new Buffer that we will use to write.
	return buffer{
		Buffer: bytes.NewBuffer(nil),
		w:      w,
	}
}

// Flush ensures that any data written to the buffer is present in the passed in
// io.Writer.
func (buf buffer) Flush() error {
	if buf.w != nil {
		if _, err := buf.w.Write(buf.Buffer.Bytes()); err != nil {
			return err
		}
		buf.Reset()
	}
	return nil
}
