package markdown_test

import (
	"bytes"
	"encoding/json"
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

var (
	updateGoldens = flag.Bool("update_goldens", false, "Set to true to update the goldens that are stored on disk.")
)

func Test(t *testing.T) {
	testModes, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatalf("Error enumerating test modes: %s", err)
	}

	for _, testMode := range testModes {
		if !testMode.IsDir() {
			continue
		}
		t.Run(testMode.Name(), func(t *testing.T) {
			testFiles, err := ioutil.ReadDir(filepath.Join("testdata/", testMode.Name()))
			if err != nil {
				t.Fatalf("Error enumerating test cases for mode %s: %s", testMode.Name(), err)
			}

			optionsFile := filepath.Join("testdata", testMode.Name(), "options.json")
			// This is the default set of options
			options := &markdown.Options{}
			optionJson, err := ioutil.ReadFile(optionsFile)
			if err == nil {
				if err := json.Unmarshal(optionJson, options); err != nil {
					t.Errorf("Unable to unmarshal %s.\nGot: %q\nError: %v", optionsFile, optionJson, err)
				}
			} else if os.IsNotExist(err) {
				// Do nothing and use the default options.
			} else {
				t.Errorf("Error opening or reading `%s`'s options.json file: %v", optionsFile, err)
			}

			for _, testFile := range testFiles {
				// Only test on input files.
				if !strings.HasSuffix(testFile.Name(), ".in.md") {
					continue
				}

				testName := testFile.Name()[0 : len(testFile.Name())-len(".in.md")]

				t.Run(testName, func(t *testing.T) {
					in, err := ioutil.ReadFile(filepath.Join("testdata", testMode.Name(), testName+".in.md"))
					if err != nil {
						t.Errorf("Unable to open input file %s.in.md: %v", testName, err)
					}

					got, err := markdown.Process("", in, options)
					if err != nil {
						t.Errorf("Error processing input: %s", err)
					}

					wantFile := filepath.Join("testdata", testMode.Name(), testName+".out.md")
					want, err := ioutil.ReadFile(wantFile)
					if err != nil {
						t.Errorf("Unable to open input file %s.in.md: %v", in, err)
					}

					d, err := diff(got, want)
					if err != nil {
						t.Errorf("Unable to diff output against golden data: %v", err)
					}

					if len(d) != 0 {
						t.Errorf("Difference of %d lines:\n%s", bytes.Count(d, []byte("\n")), string(d))
					}

					if *updateGoldens && len(d) > 0 {
						err := ioutil.WriteFile(wantFile, got, 0644)
						if err != nil {
							t.Errorf("Unable to write golden output file %s.out.md: %v", testName, err)
						}
						t.Logf("Wrote updated golden file as %s", wantFile)
					}
				})
			}
		})
	}
}

func TestLineBreak(t *testing.T) {
	input := []byte("Some text with two trailing spaces for linebreak.  \nMore      spaced      **text**      *immediately*      after      that.         \nMore than two spaces become two.\n")
	expected := []byte("Some text with two trailing spaces for linebreak.  \nMore spaced **text** *immediately* after that.  \nMore than two spaces become two.\n")

	output, err := markdown.Process("", input, nil)
	if err != nil {
		log.Fatalln(err)
	}

	diff, err := diff(expected, output)
	if err != nil {
		log.Fatalln(err)
	}

	if len(diff) != 0 {
		t.Errorf("Difference of %d lines:\n%s", bytes.Count(diff, []byte("\n")), string(diff))
	}
}

// TestDoubleSpacedListEnd tests that when the document ends with a double spaced list,
// an extra blank line isn't appended. See issue #30.
func TestDoubleSpacedListEnd(t *testing.T) {
	const reference = `-	An item.

-	Another time with a blank line in between.
`
	input := []byte(reference)
	expected := []byte(reference)

	output, err := markdown.Process("", input, nil)
	if err != nil {
		log.Fatalln(err)
	}

	diff, err := diff(expected, output)
	if err != nil {
		log.Fatalln(err)
	}

	if len(diff) != 0 {
		t.Errorf("Difference of %d lines:\n%s", bytes.Count(diff, []byte("\n")), string(diff))
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
