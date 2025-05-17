// Package cell implements the Cell Atom type and methods.
package cell

import (
	"fmt"
	"strconv"
)

// Cell is a floating-point Atom.
type Cell float64

// New returns a new Cell from a float.
func New(f float64) Cell {
	return Cell(f)
}

// Parse returns a new Cell from a string.
func Parse(s string) (Cell, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot parse cell %q - invalid format", s)
	}

	return New(f), nil
}

// Bool returns the Cell as a boolean.
func (c Cell) Bool() bool {
	return c != 0
}

// Native returns the Cell as a native value.
func (c Cell) Native() any {
	return float64(c)
}

// String returns the Cell as a string.
func (c Cell) String() string {
	return strconv.FormatFloat(float64(c), 'f', -1, 64)
}
