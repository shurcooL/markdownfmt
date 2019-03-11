An h1 header
============

Paragraphs are separated by a blank line.

2nd paragraph. *Italic*, **bold**, `monospace`. Itemized lists look like:

-	this one
-	that one
-	the other one

Nothing to note here.

> Block quotes are written like so.
>
> > They can be nested.
>
> They can span multiple paragraphs, if you like.

-	Item 1
-	Item 2
	-	Item 2a
		-	Item 2a
	-	Item 2b
-	Item 3

Hmm.

1.	Item 1
2.	Item 2
	1.	Blah.
	2.	Blah.
3.	Item 3
	-	Item 3a
	-	Item 3b

Large spacing...

1.	An entire paragraph is written here, and bigger spacing between list items is desired. This is supported too.

2.	Item 2

	1.	Blah.

	2.	Blah.

3.	Item 3

	-	Item 3a

	-	Item 3b

Last paragraph here.

An h2 header
------------

-	Paragraph right away.
-	**Big item**: Right away after header.

[Visit GitHub!](www.github.com)

![Hmm](http://example.org/image.png)

![Alt text](/path/to/img.jpg "Optional title") ![Alt text](/path/to/img.jpg "Hello \" 世界")

~~Mistaken text.~~

This (**should** be *fine*).

A \> B.

It's possible to backslash escape \<html\> tags and \`backticks\`. They are treated as text.

1986\. What a great season.

The year was 1986. What a great season.

\*literal asterisks\*.

---

http://example.com

Now a [link](www.github.com) in a paragraph. End with [link_underscore.go](www.github.com).

-	[Link](www.example.com)

### An h3 header

Here's a numbered list:

1.	first item
2.	second item
3.	third item

Note again how the actual text starts at 4 columns in (4 characters from the left side). Here's a code sample:

```
# Let me re-iterate ...
for i in 1 .. 10 { do-something(i) }
```

As you probably guessed, indented 4 spaces. By the way, instead of indenting the block, you can use delimited blocks, if you like:

```
define foobar() {
    print "Welcome to flavor country!";
}
```

(which makes copying & pasting easier). You can optionally mark the delimited block for Pandoc to syntax highlight it:

```Go
func main() {
	println("Hi.")
}
```

Here's a table.

| Name  | Age |
|-------|-----|
| Bob   | 27  |
| Alice | 23  |

Colons can be used to align columns.

| Tables        | Are           | Cool      |
|---------------|:-------------:|----------:|
| col 3 is      | right-aligned |     $1600 |
| col 2 is      |   centered!   |       $12 |
| zebra stripes |   are neat    |        $1 |
| support for   | サブタイトル  | priceless |

The outer pipes (|) are optional, and you don't need to make the raw Markdown line up prettily. You can also use inline Markdown.

| Markdown | More      | Pretty     |
|----------|-----------|------------|
| *Still*  | `renders` | **nicely** |
| 1        | 2         | 3          |

Nested Lists
============

### Codeblock within list

-	list1

	```C
	if (i == 5)
	    break;
	```

### Blockquote within list

-	list1

	> This a quote within a list.

### Table within list

-	list1

	| Header One | Header Two |
	|------------|------------|
	| Item One   | Item Two   |

### Multi-level nested

-	Item 1

	Another paragraph inside this list item is indented just like the previous paragraph.

-	Item 2

	-	Item 2a

		Things go here.

		> This a quote within a list.

		And they stay here.

	-	Item 2b

-	Item 3

Line Breaks
===========

Some text with two trailing spaces for linebreak.  
More text immediately after.  
Useful for writing poems.

Done.
