# NonoGo
[![Build Status](https://travis-ci.org/ODAEL/nonogo.svg?branch=master)](https://travis-ci.org/ODAEL/nonogo)

Nonogram solver on Golang

## Getting Started



### Installing

You can use your favourite dependency manager and add import like:

```gotemplate
import (
    "github.com/ODAEL/nonogo/holder"
)
```

## Usage Example

```go
package main

import (
	"github.com/ODAEL/nonogo/holder"
	"github.com/ODAEL/nonogo/solver"
)

func main() {

	topBox := holder.Box{Numbers: [][]int{
		{3}, {4}, {1, 4}, {9, 1}, {13}, {4, 8}, {2, 7}, {3, 8}, {2, 7}, {7}, {6}, {4}, {3},
	}}

	leftBox := holder.Box{Numbers: [][]int{
		{2, 2}, {6}, {5}, {5}, {2, 2}, {8}, {10}, {9}, {11}, {10}, {9}, {8}, {3, 4},
	}}

	nonogram := holder.BuildEmptyNonogram(13, 13)
	nonogram.TopBox = topBox
	nonogram.LeftBox = leftBox

	solver.Solve(&nonogram)

	printer := holder.CmdNonogramPrinter{Nonogram: nonogram}
	printer.PrintField()

	// Output:
	// ┏━━━━━━━━━━━━━━━━━━━━━━━━━━┓
	// ┃><><><><████><████><><><><┃
	// ┃><><><████████████><><><><┃
	// ┃><><><██████████><><><><><┃
	// ┃><██████████><><><><><><><┃
	// ┃████><████><><><><><><><><┃
	// ┃████████████████><><><><><┃
	// ┃████████████████████><><><┃
	// ┃><><██████████████████><><┃
	// ┃><><██████████████████████┃
	// ┃><><><████████████████████┃
	// ┃><><><><██████████████████┃
	// ┃><><><><████████████████><┃
	// ┃><><><██████><████████><><┃
	// ┗━━━━━━━━━━━━━━━━━━━━━━━━━━┛
}
```

## Running the tests

Each package contains a tests which can be run with `go test ./...`

## Dependencies

* [stretchr/testify/assert](https://github.com/stretchr/testify/assert) - Library for asserts in tests

## Authors

* **Aleksander Penko** - *Initial work* - [ODAEL](https://github.com/ODAEL)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
