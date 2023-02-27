# multistrings

## What is multistrings?

**multistrings** is a small library to manipulate multiple strings as a whole.

## Usage

```golang
import (
    "github.com/tpfeiffer67/multistrings"
)
```

## Functions

- New creates a MultiStrings with a given capacity.
- NewWithGivenLen creates a MultiStrings with a given length.
- NewFromStringSlice creates a MultiStrings from a string slice.
- NewWithPresetContent creates a MultiStrings with a preset line count and content.
- NewMultiStringsFromString creates a MultiStrings from a string.
- WriteToFile writes the MultiStrings to a file.
- ToString concatenates the MultiStrings into a string with a given line separator.
- String, the Stringer interface.
- Width returns the length of the longest string in the MultiStrings.
- Cut truncates the MultiStrings to a given maximum width by cutting strings that exceed it.
- Limit truncates the MultiStrings to a given maximum width by splitting strings that exceed it.
- PadRight pads the MultiStrings on the right side of each string with a fill pattern to achieve a given width.
- Sort sorts the MultiStrings in lexicographically increasing order.

## Dependencies
The package [runesstr](https://github.com/tpfeiffer67/runesstr) is imported to work with Unicode characters in the strings.

## Examples

Example 1
```go
package main

import (
	"fmt"

	"github.com/tpfeiffer67/multistrings"
)

const loremipsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit,
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

func main() {
	ms := multistrings.NewFromString(loremipsum)
	fmt.Println(ms)
}
```

## Contributing

If you'd like to contribute to this repository, feel free to submit a pull request. Please include tests for any new functions or changes to existing functions.

## License

This repository is licensed under the MIT License. See the LICENSE file for more information.


## Author

- [@TPfeiffer67](https://www.github.com/TPfeiffer67)
