package ansi

import "testing"

func TestConstruct(t *testing.T) {
	t.Log(Construct(ReverseVid, ColorGreen), "Hello World!", Clear)
	t.Log(Construct(Bold, Underline, BackBlue, ColorWhite), "Hello World!", Clear)
}
