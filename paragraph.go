// The paragraph library provides a set of utility functions to work with a slice of strings.
// These functions can be used to manipulate and transform a slice of strings, such as creating a new slice with a given capacity,
// splitting a string into a slice of strings, writing a slice of strings to a file, limiting the width of strings in a slice,
// padding strings in a slice with a fill pattern, sorting a slice of strings, and appending one slice of strings to another.

// The paragraph library provides a Paragraph type, which is an alias for a slice of strings.
// The constant MultiStringsMaxWidth sets a arbitrary limit to the maximum width of a string in the slice.

package paragraph

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/tpfeiffer67/runesstr"
)

const MultiStringsMaxWidth = 1000

type Paragraph []string

// New creates and returns a new Paragraph slice with a given capacity.
// - cap is an integer representing the desired capacity of the slice.
func New(cap int) Paragraph {
	if cap < 0 {
		cap = 1
	}
	return (make([]string, 0, cap))
}

// NewWithGivenLen creates and returns a new Paragraph slice with a given length.
// - l is the desired length of the slice.
func NewWithGivenLen(length int) Paragraph {
	if length < 0 {
		length = 1
	}
	return (make([]string, length, length))
}

// NewFromStringSlice creates and returns a new Paragraph slice from a string slice.
// - s is a slice of strings from which to create the new Paragraph slice.
func NewFromStringSlice(s []string) (linesOut Paragraph) {
	linesOut = NewWithGivenLen(len(s))
	for index, l := range s {
		linesOut[index] = l
	}
	return
}

// NewWithPresetContent creates and returns a new Paragraph slice with a preset line count and content.
// - line is a string representing the content to be added to each line of the new slice.
// - count is the desired number of lines in the new slice.
func NewWithPresetContent(line string, count int) (linesOut Paragraph) {
	linesOut = NewWithGivenLen(count)
	for i := 0; i < count; i++ {
		linesOut[i] = line
	}
	return
}

// NewFromString creates and returns a new Paragraph slice from a string.
// The given string is split into separate substrings at each occurrence of a newline character.
// - s is a string from which to create the new Paragraph slice.
func NewFromString(s string) Paragraph {
	return strings.Split(s, "\n")
}

// WriteToFile writes the Paragraph slice to a file with a given filename.
// - fileName is the name of the file to which to write the slice.
func (lines Paragraph) WriteToFile(fileName string) (err error) {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Unable to create file with given file name: %w", err)
	}
	for _, s := range lines {
		_, err = f.WriteString(s + "\n")
		if err != nil {
			return fmt.Errorf("Unable to write string to file: %w", err)
		}
	}
	f.Sync()
	return
}

// ToString concatenates the Paragraph slice into a string with a given line separator.
// - linesSeparator is a string representing the separator to use between each line.
func (lines Paragraph) ToString(linesSeparator string) (result string) {
	for _, s := range lines {
		result = result + s + linesSeparator
	}
	return
}

// String implements the Stringer interface for Paragraph.
func (lines Paragraph) String() string {
	return lines.ToString("\n")
}

// Width returns the width of the Paragraph slice, which is the length of the longest string in the slice.
func (lines Paragraph) Width() (width int) {
	for _, s := range lines {
		width = maxint(runesstr.Length(s), width)
	}
	return
}

// Cut truncates the Paragraph slice to a given maximum width by cutting strings that exceed it.
// - maxWidth is the maximum width to which to truncate the strings.
func (linesIn Paragraph) Cut(maxWidth int) (linesOut Paragraph) {
	if maxWidth < 1 {
		linesOut = linesIn
		return
	}
	linesOut = New(len(linesIn))
	for _, s := range linesIn {
		if runesstr.Length(s) > maxWidth {
			linesOut = append(linesOut, runesstr.Left(s, maxWidth))
		} else {
			linesOut = append(linesOut, s)
		}
	}
	return
}

// Limit truncates the Paragraph slice to a given maximum width by splitting strings that exceed it.
// - maxWidth is the maximum width to which to truncate the strings.
func (linesIn Paragraph) Limit(maxWidth int) (linesOut Paragraph) {
	if maxWidth < 1 {
		return linesIn
	}
	linesOut = New(len(linesIn)) // at least the same len than linesIn
	for _, s := range linesIn {
		for {
			var sl string
			if runesstr.Length(s) > maxWidth {
				sl, s = runesstr.SplitOnNearestSpace(s, maxWidth)
				linesOut = append(linesOut, sl)
			} else {
				linesOut = append(linesOut, s)
				break
			}
		}
	}
	return
}

// PadRight pads the Paragraph slice on the right side with a given fill pattern to a given width.
// - fillPattern represents the pattern to use for padding.
// - width is the desired width of each line after padding.
func (linesIn Paragraph) PadRight(fillPattern string, width int) (linesOut Paragraph) {
	if runesstr.Length(fillPattern) == 0 {
		return linesIn
	}
	if width < 1 || width > MultiStringsMaxWidth { // Here we set a limit to width
		return linesIn
	}
	l := len(linesIn)
	linesOut = NewWithGivenLen(l)
	for i := 0; i < l; i++ {
		linesOut[i] = runesstr.PadRight(linesIn[i], fillPattern, width)
	}
	return
}

// Surround surrounds each line of the Paragraph slice with a given left and right string.
// - left is the left-side surround for each line.
// - right is the right-side surround.
func (linesIn Paragraph) Surround(left string, right string) (linesOut Paragraph) {
	l := len(linesIn)
	linesOut = NewWithGivenLen(l)
	for i := 0; i < l; i++ {
		linesOut[i] = left + linesIn[i] + right
	}
	return
}

// Sort sorts the Paragraph slice in lexicographic order and returns the sorted slice.
func (linesIn Paragraph) Sort() Paragraph {
	sort.Strings(linesIn)
	return linesIn
}

// Append appends the strings from another Paragraph slice to the end of the current one.
func (linesIn Paragraph) Append(linesToAppend Paragraph) Paragraph {
	for _, s := range linesToAppend {
		linesIn = append(linesIn, s)
	}
	return linesIn
}
