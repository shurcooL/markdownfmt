// This package provides a debug renderer. It prints the method names and parameters as they're called.
package debug

import (
	"bytes"
	"fmt"

	. "gist.github.com/6418290.git"

	"github.com/shurcooL/go-goon"
)

type debugRenderer struct {
}

func NewRenderer() *debugRenderer {
	return &debugRenderer{}
}

func (_ *debugRenderer) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(text), lang)
}
func (_ *debugRenderer) BlockQuote(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(text))
}
func (_ *debugRenderer) BlockHtml(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(text))
}
func (_ *debugRenderer) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(text(), level, id)
}
func (_ *debugRenderer) HRule(out *bytes.Buffer) {
	fmt.Println(GetParentFuncAsString())
}
func (_ *debugRenderer) List(out *bytes.Buffer, text func() bool, flags int) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(text(), flags)
}
func (_ *debugRenderer) ListItem(out *bytes.Buffer, text []byte, flags int) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(text), flags)
}
func (_ *debugRenderer) Paragraph(out *bytes.Buffer, text func() bool) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(text())
}
func (_ *debugRenderer) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(header), string(body), columnData)
}
func (_ *debugRenderer) TableRow(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(text))
}
func (_ *debugRenderer) TableHeaderCell(out *bytes.Buffer, text []byte, flags int) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(text), flags)
}
func (_ *debugRenderer) TableCell(out *bytes.Buffer, text []byte, flags int) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(text), flags)
}
func (_ *debugRenderer) Footnotes(out *bytes.Buffer, text func() bool) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(text())
}
func (_ *debugRenderer) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(name), string(text), flags)
}

func (_ *debugRenderer) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(link), kind)
}
func (_ *debugRenderer) CodeSpan(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(text))
}
func (_ *debugRenderer) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(text))
}
func (_ *debugRenderer) Emphasis(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(text))
}
func (_ *debugRenderer) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(link), string(title), string(alt))
}
func (_ *debugRenderer) LineBreak(out *bytes.Buffer) {
	fmt.Println(GetParentFuncAsString())
}
func (_ *debugRenderer) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(link), string(title), string(content))
}
func (_ *debugRenderer) RawHtmlTag(out *bytes.Buffer, tag []byte) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(tag))
}
func (_ *debugRenderer) TripleEmphasis(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(text))
}
func (_ *debugRenderer) StrikeThrough(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(text))
}
func (_ *debugRenderer) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(ref), id)
}

func (_ *debugRenderer) Entity(out *bytes.Buffer, entity []byte) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(entity))
}
func (_ *debugRenderer) NormalText(out *bytes.Buffer, text []byte) {
	fmt.Println(GetParentFuncAsString())
	goon.DumpExpr(string(text))
}

func (_ *debugRenderer) DocumentHeader(out *bytes.Buffer) {
	fmt.Println(GetParentFuncAsString())
}
func (_ *debugRenderer) DocumentFooter(out *bytes.Buffer) {
	fmt.Println(GetParentFuncAsString())
}

func (_ *debugRenderer) GetFlags() int {
	fmt.Println(GetParentFuncAsString())
	return 0
}
