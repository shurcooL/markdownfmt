module github.com/shurcooL/markdownfmt

go 1.21

require (
	github.com/mattn/go-runewidth v0.0.13
	github.com/shurcooL/go v0.0.0-20200502201357-93f07166e636
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
)

// 1.5.2 and 1.6.0 have a bug with code block + list parsing where code blocks after a list get consumed into the final list item (which is the opposite of the bug they were trying to fix when they broke it: https://github.com/russross/blackfriday/issues/239, https://github.com/russross/blackfriday/issues/495, https://github.com/russross/blackfriday/issues/485, https://github.com/russross/blackfriday/pull/521
require github.com/russross/blackfriday v1.5.1

require (
	github.com/rivo/uniseg v0.2.0 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
)
