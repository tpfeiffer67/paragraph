package multistrings

import (
	"github.com/tpfeiffer67/runesstr"
)

type BoxPattern struct {
	TopLeftCorner     string
	TopBorder         string
	TopRightCorner    string
	LeftBorder        string
	RightBorder       string
	BottomLeftCorner  string
	BottomBorder      string
	BottomRightCorner string
}

// BoxSettings{30, "", Left, "", Left}
type BoxSettings struct {
	Width            int
	TopLabel         string
	TopLabelAlign    LabelAlign
	BottomLabel      string
	BottomLabelAlign LabelAlign
}

var boxPatterns = [BoxStyleCount]BoxPattern{
	{"", "", "", "", "", "", "", ""},
	{" ", " ", " ", " ", " ", " ", " ", " "},

	{"┌", "─", "┐", "│", "│", "└", "─", "┘"},
	{"╭", "─", "╮", "│", "│", "╰", "─", "╯"},
	{"┏", "━", "┓", "┃", "┃", "┗", "━", "┛"},
	{"╒", "═", "╕", "│", "│", "╘", "═", "╛"},
	{"╓", "─", "╖", "║", "║", "╙", "─", "╜"},
	{"╔", "═", "╗", "║", "║", "╚", "═", "╝"},
	{"▛", "▀", "▜", "▌", "▐", "▙", "▄", "▟"},
	{"▞", "▀", "▚", "▌", "▐", "▚", "▄", "▞"},
	{"█", "▀", "█", "█", "█", "█", "▄", "█"},
	{"░", "░", "░", "░", "░", "░", "░", "░"},
	{"▒", "▒", "▒", "▒", "▒", "▒", "▒", "▒"},
	{"▓", "▓", "▓", "▓", "▓", "▓", "▓", "▓"},
	{"█", "█", "█", "█", "█", "█", "█", "█"},

	{".", ".", ".", ":", ":", ":", ".", ":"},
	{"◆", "◆", "◆", "◆", "◆", "◆", "◆", "◆"},
	{"╭", "╼", "╮", "╽", "╿", "╰", "╾", "╯"},
	{`╱`, "▔", "╲", "│", "│", "╲", "▁", `╱`},
	{"▁▂▃", "▃", "▃▂▁", "▌", "▐", "▜▃▂▁", "▁", "▁▂▃▛"},
	{"", "▁▂▃▂", "", "█", "█", "█", "▃▂▁▂", "█"},
}

func GetBoxPattern(style BoxStyle) BoxPattern {
	if style < 0 || style > BoxStyleLastValue {
		return boxPatterns[BoxStyleSingleLine]
	}
	return boxPatterns[style]
}

func (linesIn MultiStrings) AutoBox(settings BoxSettings, pattern BoxPattern) MultiStrings {
	if settings.Width < 1 || settings.Width > MultiStringsMaxWidth || pattern == boxPatterns[BoxStyleNone] {
		return linesIn
	}
	w := linesIn.Width()
	/*
		// Top label must not be truncated
		l := len(settings.TopLabel)
		if l > w {
			w = l
		}*/
	settings.Width = w
	return linesIn.PadRight(" ", w).Box(settings, pattern)
}

func (linesIn MultiStrings) Box(settings BoxSettings, pattern BoxPattern) (linesOut MultiStrings) {
	if settings.Width < 1 || settings.Width > MultiStringsMaxWidth || pattern == boxPatterns[BoxStyleNone] {
		return linesIn
	}
	width := settings.Width
	bordersWidth := runesstr.Length(pattern.LeftBorder) + runesstr.Length(pattern.RightBorder)
	toplabel, ltopleft, ltopright := processLabel(settings.TopLabel, settings.TopLabelAlign, width, bordersWidth, runesstr.Length(pattern.TopLeftCorner)+runesstr.Length(pattern.TopRightCorner))
	bottomlabel, lbottomleft, lbottomright := processLabel(settings.BottomLabel, settings.BottomLabelAlign, width, bordersWidth, runesstr.Length(pattern.BottomLeftCorner)+runesstr.Length(pattern.BottomRightCorner))

	l := len(linesIn)
	linesOut = NewWithGivenLen(l + 2)
	linesOut[0] = pattern.TopLeftCorner + runesstr.PadRight("", pattern.TopBorder, ltopleft) + toplabel + runesstr.PadRight("", pattern.TopBorder, ltopright) + pattern.TopRightCorner
	for i := 0; i < l; i++ {
		linesOut[i+1] = pattern.LeftBorder + linesIn[i] + pattern.RightBorder
	}
	linesOut[l+1] = pattern.BottomLeftCorner + runesstr.PadRight("", pattern.BottomBorder, lbottomleft) + bottomlabel + runesstr.PadRight("", pattern.BottomBorder, lbottomright) + pattern.BottomRightCorner
	return
}

func processLabel(label string, align LabelAlign, width int, bordersWidth int, cornersWidth int) (rLabel string, lleft int, lright int) {
	l := width + bordersWidth - runesstr.Length(label) - cornersWidth
	if l <= 0 {
		l = 0
		rLabel = runesstr.Left(label, width)
	} else {
		rLabel = label
	}

	switch align {
	case LabelAlignRight:
		lleft = l
		lright = 0
	case LabelAlignCenter:
		lleft = l/2 + l%2
		lright = l / 2
	default: // LabelAlignLeft
		lleft = 0
		lright = l
	}

	return
}
