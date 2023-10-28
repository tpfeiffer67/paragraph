# paragraph

## What is paragraph?

**paragraph** is a small library to manipulate multiple strings as a whole.

## Usage

```golang
import (
    "github.com/tpfeiffer67/paragraph"
)
```

## Functions

- New creates a Paragraph with a given capacity.
- NewWithGivenLen creates a Paragraph with a given length.
- NewFromStringSlice creates a Paragraph from a string slice.
- NewWithPresetContent creates a Paragraph with a preset line count and content.
- NewFromString creates a Paragraph from a string.
- WriteToFile writes the Paragraph to a file.
- ToString concatenates the Paragraph into a string with a given line separator.
- String, the Stringer interface.
- Width returns the length of the longest string in the Paragraph.
- Cut truncates the Paragraph to a given maximum width by cutting strings that exceed it.
- Limit truncates the Paragraph to a given maximum width by splitting strings that exceed it.
- PadRight pads the Paragraph on the right side of each string with a fill pattern to achieve a given width.
- Sort sorts the Paragraph in lexicographically increasing order.

## Dependencies
The package [runesstr](https://github.com/tpfeiffer67/runesstr) is imported to work with Unicode characters in the strings.

## Examples

Example 1
```go
package main

import (
	"fmt"

	"github.com/tpfeiffer67/paragraph"
)

const loremipsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit,
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

func main() {
	ms := paragraph.NewFromString(loremipsum)
	fmt.Println(ms)
}
```

## License

This repository is licensed under the MIT License. See the LICENSE file for more information.


## Author

- [@TPfeiffer67](https://www.github.com/TPfeiffer67)
