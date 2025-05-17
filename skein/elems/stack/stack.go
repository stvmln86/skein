// Package stack implements the Stack type and methods.
package stack

import (
	"fmt"
	"strings"

	"github.com/stvmln86/skein/skein/atoms/cell"
)

// Stack is a last-in-first-out stack of Cells.
type Stack struct {
	Cells []cell.Cell
}

// New returns a new Stack from zero or more Cells.
func New(cs ...cell.Cell) *Stack {
	return &Stack{cs}
}

// Empty returns true if there are no Cells on the Stack.
func (st *Stack) Empty() bool {
	return len(st.Cells) == 0
}

// Len returns the number of Cells on the Stack.
func (st *Stack) Len() int {
	return len(st.Cells)
}

// Pop removes and returns the top Cell on the Stack.
func (st *Stack) Pop() (cell.Cell, error) {
	if st.Empty() {
		return 0, fmt.Errorf("cannot pop stack - is empty")
	}

	c := st.Cells[len(st.Cells)-1]
	st.Cells = st.Cells[:len(st.Cells)-1]
	return c, nil
}

// PopSlice removes and returns the top N Cells on the Stack.
func (st *Stack) PopSlice(n int) ([]cell.Cell, error) {
	if len(st.Cells) < n {
		return nil, fmt.Errorf("cannot pop stack - is insufficient")
	}

	cs := st.Cells[len(st.Cells)-n:]
	st.Cells = st.Cells[:len(st.Cells)-n]
	return cs, nil
}

// Push appends a Cell to the top of the Stack.
func (st *Stack) Push(c cell.Cell) {
	st.Cells = append(st.Cells, c)
}

// PushSlice appends a Cell slice to the top of the Stack.
func (st *Stack) PushSlice(cs []cell.Cell) {
	st.Cells = append(st.Cells, cs...)
}

// String returns the Stack as a string.
func (st *Stack) String() string {
	var ss []string
	for _, c := range st.Cells {
		ss = append(ss, c.String())
	}

	return strings.Join(ss, " ")
}
