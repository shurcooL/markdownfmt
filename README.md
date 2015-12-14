markdownfmt [![Build Status](https://travis-ci.org/shurcooL/markdownfmt.svg?branch=master)](https://travis-ci.org/shurcooL/markdownfmt) [![GoDoc](https://godoc.org/github.com/shurcooL/markdownfmt?status.svg)](https://godoc.org/github.com/shurcooL/markdownfmt)
===================================================================================================================================================================================================================================================================

Like `gofmt`, but for Markdown.

![Markdown Format Demo](https://github.com/shurcooL/atom-markdown-format/blob/master/Demo.gif?raw=true)

Installation
------------

```bash
go get -u github.com/shurcooL/markdownfmt
```

Add `$GOPATH/bin` to your `$PATH` or copy `$GOPATH/bin/markdownfmt` to your `$PATH`.

Usage
-----

```
usage: markdownfmt [flags] [path ...]
  -d=false: display diffs instead of rewriting files
  -l=false: list files whose formatting differs from markdownfmt's
  -w=false: write result to (source) file instead of stdout
```

Editor Plugins
--------------

-	[markdown-format](https://atom.io/packages/markdown-format) for Atom.
-	[vim-markdownfmt](https://github.com/moorereason/vim-markdownfmt) for Vim.
-	Built-in in Conception.
-	Add a plugin for your favorite editor here?

License
-------

-	[MIT License](http://opensource.org/licenses/mit-license.php)
