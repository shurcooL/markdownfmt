markdownfmt
===========

Like `gofmt`, but for Markdown.

![Markdown Format Demo](https://github.com/shurcooL/atom-markdown-format/blob/master/Demo.gif?raw=true)

Installation
------------

```bash
$ go get -u github.com/shurcooL/markdownfmt
```

Add `$GOPATH/bin` to your `$PATH` or copy `$GOPATH/bin/markdownfmt` to your `$PATH`.

### Don't have Go on OSX?

Install it via Homebrew:

```bash
brew install go mercurial
go get -u github.com/shurcooL/markdownfmt
mkdir ~/gocode
```

Add to bash profile (`~/.bash_profile` or `~/.profile`):

```bash
export GOPATH="$HOME/gocode"
export PATH="$PATH:$GOPATH/bin"
```

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

- [markdown-format](https://atom.io/packages/markdown-format) for Atom.
- Built-in in Conception.
- Add a plugin for your favorite editor here?
