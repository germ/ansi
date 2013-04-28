// A metric fuckton of ANSI escape codes
// Most of these codes are garbage but I need the colors
// And a few more. The code is operator agnostic
package ansi

import "strconv"

type AnsiCode int

const Clear = "\x1b[0m"
const (
	Reset AnsiCode = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkFast
	ReverseVid
	Conceal
	CrossOut
	PrimFont
	OneAltFont
	TwoAltFont
	ThreeAltFont
	FourAltFont
	FiveAltFont
	SixAltFont
	SevenAltFont
	EightAltFont
	NineAltFont
	Fraktur
	BoldOffDblUndr
	Normal
	NotItalic
	UnderlineNone
	BlinkOff
	Reversed
	ImagePositive
	Reveal
	NotCrossedOut
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
	XTerm256
	DefaultColor
	BackBlack
	BackRed
	BackGreen
	BackYellow
	BackBlue
	BackMagenta
	BackCyan
	BackWhite
	XTermBack256
	DefaultBack
	Reserved
	Framed
	Encircled
	Overlined
	NotFrameEncircle
	NotOverlined
	Reserved1
	Reserved2
	Reserved3
	Reserved4
	IdeoUnderline
	IdeoDblUnderline
	IdeoOverline
	IdeoDblOverline
	IdeoStress
)

// Returns a string embeded with the requested codes
func Construct(ops ...AnsiCode) string {
	res := "\x1b["
	for i, v := range ops {
		res += strconv.Itoa(int(v))

		if i != len(ops)-1 {
			res += ";"
		}
	}
	res += "m"

	return res
}
