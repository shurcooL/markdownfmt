package markdown

import (
	"bytes"
	"testing"

	"github.com/mattn/go-runewidth"
)

// Test that string "**bold**" has width 8 when terminal is off, and no ANSI escape codes are added.
func TestNormalStringWidth(t *testing.T) {
	r := NewRenderer(nil).(*markdownRenderer)
	var buf bytes.Buffer
	r.Strong(&buf, true)
	buf.WriteString("bold")
	r.Strong(&buf, false)
	if got, want := buf.String(), "**bold**"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	if got, want := r.stringWidth("**bold**"), len("**bold**"); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := r.stringWidth("\x1b[1m**bold**\x1b[0m"), runewidth.StringWidth("\x1b[1m**bold**\x1b[0m"); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Test that string "\x1b[1m**bold**\x1b[0m" has width 8 when terminal is on. ANSI escape codes should not count for string width.
func TestTerminalStringWidth(t *testing.T) {
	r := NewRenderer(&Options{Terminal: true}).(*markdownRenderer)
	var buf bytes.Buffer
	r.Strong(&buf, true)
	buf.WriteString("bold")
	r.Strong(&buf, false)
	if got, want := buf.String(), "\x1b[1m**bold**\x1b[0m"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	if got, want := r.stringWidth("\x1b[1m**bold**\x1b[0m"), len("**bold**"); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, dontWant := r.stringWidth("\x1b[1m**bold**\x1b[0m"), runewidth.StringWidth("\x1b[1m**bold**\x1b[0m"); got == dontWant {
		t.Errorf("got %q, dontWant %q", got, dontWant)
	}
}
