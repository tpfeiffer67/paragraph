# multistrings

## What is multistrings?

**multistrings** is a small library to manipulate multiple strings as a whole.

## Usage

```golang
import (
    "github.com/tpfeiffer67/multistrings"
)
```

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
	ms := multistrings.NewMultiStringsFromString(loremipsum)
	fmt.Println(ms)
}
```

## Contributing

If you'd like to contribute to this repository, feel free to submit a pull request. Please include tests for any new functions or changes to existing functions.

## License

This repository is licensed under the MIT License. See the LICENSE file for more information.


## Author

- [@TPfeiffer67](https://www.github.com/TPfeiffer67)
