package paragraph

// IMPORTANT: This file was auto-generated by goenum.exe and should not be modified directly.
// Any changes made to this file will be overwritten the next time goenum.exe is run.
// This file was generated based on the original description file located at ./goenum/AccoladesStyle.goenum.
// The template used to generate this file can be found at ./goenum/goenum.template.
// To make changes to the enumeration, please update the original description file and re-run goenum.exe.
// The source code for goenum can be found here https://github.com/tpfeiffer67/goenum

import (
	"errors"
	"strings"
)

type AccoladesStyle int

const (
	AccoladesStyleCount     = 3
	AccoladesStyleMaxIndex  = int(AccoladesStyleUnicode)
	AccoladesStyleLastValue = AccoladesStyleUnicode
)

const (
	AccoladesStyleNone AccoladesStyle = iota
	AccoladesStyleAscii
	AccoladesStyleUnicode
)

func (v AccoladesStyle) String() string {
	return [...]string{
		"AccoladesStyleNone",
		"AccoladesStyleAscii",
		"AccoladesStyleUnicode",
	}[v]
}

func AccoladesStyleFromString(s string) (AccoladesStyle, error) {
	var suffix string
	if strings.HasPrefix(s, "AccoladesStyle") {
		l := len("AccoladesStyle")
		if l < len(s) {
			suffix = s[l:]
		}
	} else {
		suffix = s
	}
	switch suffix {
	case "None":
		return AccoladesStyleNone, nil
	case "Ascii":
		return AccoladesStyleAscii, nil
	case "Unicode":
		return AccoladesStyleUnicode, nil
	}
	return AccoladesStyle(0), errors.New("String does not correspond to any existing AccoladesStyle values")
}
