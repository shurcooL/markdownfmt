// This package provides a debug renderer. It prints the method names and parameters as they're called.
package debug

import (
	"bytes"
	"fmt"

	"github.com/russross/blackfriday"
	. "github.com/shurcooL/go/gists/gist6418290"
)

type debugRenderer struct {
	real blackfriday.Renderer
}

func NewRenderer(realRenderer blackfriday.Renderer) *debugRenderer {
	return &debugRenderer{real: realRenderer}
}

func (dr *debugRenderer) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	fmt.Println(GetParentFuncArgsAsString(string(text), lang))
	dr.real.BlockCode(out, text, lang)
}
func (dr *debugRenderer) BlockQuote(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(text)))
	dr.real.BlockQuote(out, text)
}
func (_ *debugRenderer) BlockHtml(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(text)))
}
func (_ *debugRenderer) TitleBlock(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(text)))
}
func (_ *debugRenderer) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	fmt.Println(GetParentFuncArgsAsString(level, id))
}
func (_ *debugRenderer) HRule(out *bytes.Buffer) {
	fmt.Println(GetParentFuncArgsAsString())
}
func (dr *debugRenderer) List(out *bytes.Buffer, text func() bool, flags int) {
	fmt.Println(GetParentFuncArgsAsString(flags))
	debugText := func() bool {
		b := text()
		fmt.Println("text", GetParentFuncArgsAsString(), "->", b)
		return b
	}
	dr.real.List(out, debugText, flags)
}
func (dr *debugRenderer) ListItem(out *bytes.Buffer, text []byte, flags int) {
	fmt.Println(GetParentFuncArgsAsString(string(text), flags))
	dr.real.ListItem(out, text, flags)
}
func (dr *debugRenderer) Paragraph(out *bytes.Buffer, text func() bool) {
	fmt.Println(GetParentFuncArgsAsString())
	debugText := func() bool {
		b := text()
		fmt.Println("text", GetParentFuncArgsAsString(), "->", b)
		return b
	}
	dr.real.Paragraph(out, debugText)
}
func (_ *debugRenderer) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	fmt.Println(GetParentFuncArgsAsString(string(header), string(body), columnData))
}
func (_ *debugRenderer) TableRow(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(text)))
}
func (_ *debugRenderer) TableHeaderCell(out *bytes.Buffer, text []byte, align int) {
	fmt.Println(GetParentFuncArgsAsString(string(text), align))
}
func (_ *debugRenderer) TableCell(out *bytes.Buffer, text []byte, align int) {
	fmt.Println(GetParentFuncArgsAsString(string(text), align))
}
func (_ *debugRenderer) Footnotes(out *bytes.Buffer, text func() bool) {
	fmt.Println(GetParentFuncArgsAsString())
}
func (_ *debugRenderer) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	fmt.Println(GetParentFuncArgsAsString(string(name), string(text), flags))
}

func (_ *debugRenderer) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	fmt.Println(GetParentFuncArgsAsString(string(link), kind))
}
func (_ *debugRenderer) CodeSpan(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(text)))
}
func (_ *debugRenderer) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(text)))
}
func (_ *debugRenderer) Emphasis(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(text)))
}
func (_ *debugRenderer) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(link), string(title), string(alt)))
}
func (dr *debugRenderer) LineBreak(out *bytes.Buffer) {
	fmt.Println(GetParentFuncArgsAsString())
	dr.real.LineBreak(out)
}
func (_ *debugRenderer) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(link), string(title), string(content)))
}
func (_ *debugRenderer) RawHtmlTag(out *bytes.Buffer, tag []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(tag)))
}
func (_ *debugRenderer) TripleEmphasis(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(text)))
}
func (_ *debugRenderer) StrikeThrough(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(text)))
}
func (_ *debugRenderer) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	fmt.Println(GetParentFuncArgsAsString(string(ref), id))
}

func (dr *debugRenderer) Entity(out *bytes.Buffer, entity []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(entity)))
	dr.real.Entity(out, entity)
}
func (dr *debugRenderer) NormalText(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncArgsAsString(string(text)))
	dr.real.NormalText(out, text)
}

func (_ *debugRenderer) DocumentHeader(out *bytes.Buffer) {
	fmt.Println(GetParentFuncArgsAsString())
}
func (_ *debugRenderer) DocumentFooter(out *bytes.Buffer) {
	fmt.Println(GetParentFuncArgsAsString())
}

func (dr *debugRenderer) GetFlags() int {
	fmt.Println(GetParentFuncArgsAsString())
	return dr.real.GetFlags()
}
