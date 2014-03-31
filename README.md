markdownfmt
===========

`markdownfmt` is to Markdown, like `gofmt` is to Go source code.

Installation
------------

```bash
$ go get -u github.com/shurcooL/markdownfmt
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
