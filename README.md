markdownfmt
===========

[![Go Reference](https://pkg.go.dev/badge/github.com/shurcooL/markdownfmt.svg)](https://pkg.go.dev/github.com/shurcooL/markdownfmt)

Like `gofmt`, but for Markdown.

![Markdown Format Demo](https://github.com/shurcooL/atom-markdown-format/blob/master/Demo.gif?raw=true)

Note that `markdownfmt` works with pure Markdown files. If you want to use it with Markdown files that have front matter, consider one of [alternatives](#alternatives) that supports that.

Installation
------------

```sh
go install github.com/shurcooL/markdownfmt@latest
```

Usage
-----

```sh
usage: markdownfmt [flags] [path ...]
  -d  display diffs instead of rewriting files
  -l  list files whose formatting differs from markdownfmt's
  -w  write result to (source) file instead of stdout
```

Editor Plugins
--------------

-	[vim-markdownfmt](https://github.com/moorereason/vim-markdownfmt) for Vim.
-	[emacs-markdownfmt](https://github.com/nlamirault/emacs-markdownfmt) for Emacs.
-	[vscode-markdownfmt](https://marketplace.visualstudio.com/itemdetails?itemName=AnmolSinghJaggi.vscode-markdownfmt) for Visual Studio Code.
-	Built-in in Conception.
-	[markdown-format](https://atom.io/packages/markdown-format) for Atom (deprecated).
-	Add a plugin for your favorite editor here?

Alternatives
------------

-	[`mdfmt`](https://github.com/moorereason/mdfmt) - Fork of `markdownfmt` that adds front matter support.
-	[`tidy-markdown`](https://github.com/slang800/tidy-markdown) - Project with similar goals, but written in JS and based on a slightly different [styleguide](https://github.com/slang800/markdown-styleguide).
-	[Flowmark](https://github.com/jlevy/atom-flowmark) - A JS-based Atom plugin with line wrapping, YAML frontmatter support, and other normalization features.

License
-------

-	[MIT License](LICENSE)
