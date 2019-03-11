package markdown_test

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/shurcooL/markdownfmt/markdown"
)

func Example() {
	input := []byte(`Title
=

This is a new paragraph. I wonder    if I have too     many spaces.
What about new paragraph.
But the next one...

  Is really new.

1. Item one.
1. Item TWO.


Final paragraph.
`)

	output, err := markdown.Process("", input, nil)
	if err != nil {
		log.Fatalln(err)
	}

	os.Stdout.Write(output)

	// Output:
	// Title
	// =====
	//
	// This is a new paragraph. I wonder if I have too many spaces. What about new paragraph. But the next one...
	//
	// Is really new.
	//
	// 1.	Item one.
	// 2.	Item TWO.
	//
	// Final paragraph.
	//
}

func Example_two() {
	input := []byte(`Title
==

Subtitle
---

How about ` + "`this`" + ` and other stuff like *italic*, **bold** and ***super    extra***.
`)

	output, err := markdown.Process("", input, nil)
	if err != nil {
		log.Fatalln(err)
	}

	os.Stdout.Write(output)

	// Output:
	// Title
	// =====
	//
	// Subtitle
	// --------
	//
	// How about `this` and other stuff like *italic*, **bold** and ***super extra***.
	//
}

var updateFlag = flag.Bool("update", false, "Update golden files.")

func Test(t *testing.T) {
	fis, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	for _, fi := range fis {
		if !strings.HasSuffix(fi.Name(), ".in.md") {
			continue
		}
		name := strings.TrimSuffix(fi.Name(), ".in.md")
		t.Run(name, func(t *testing.T) {
			got, err := markdown.Process(filepath.Join("testdata", fi.Name()), nil, nil)
			if err != nil {
				t.Fatal("markdown.Process:", err)
			}
			if *updateFlag {
				err := ioutil.WriteFile(filepath.Join("testdata", name+".golden.md"), got, 0644)
				if err != nil {
					t.Fatal(err)
				}
				return
			}

			want, err := ioutil.ReadFile(filepath.Join("testdata", name+".golden.md"))
			if err != nil {
				t.Fatal(err)
			}

			diff, err := diff(got, want)
			if err != nil {
				t.Fatal(err)
			}
			if len(diff) != 0 {
				t.Errorf("difference of %d lines:\n%s", bytes.Count(diff, []byte("\n")), string(diff))
			}
		})
	}
}

// TODO: Factor out.
func diff(b1, b2 []byte) (data []byte, err error) {
	f1, err := ioutil.TempFile("", "markdownfmt")
	if err != nil {
		return
	}
	defer os.Remove(f1.Name())
	defer f1.Close()

	f2, err := ioutil.TempFile("", "markdownfmt")
	if err != nil {
		return
	}
	defer os.Remove(f2.Name())
	defer f2.Close()

	f1.Write(b1)
	f2.Write(b2)

	data, err = exec.Command("diff", "-u", f1.Name(), f2.Name()).CombinedOutput()
	if len(data) > 0 {
		// diff exits with a non-zero status when the files don't match.
		// Ignore that failure as long as we get output.
		err = nil
	}
	return
}
