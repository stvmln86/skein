package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/skein/skein/atoms/cell"
)

func TestNew(t *testing.T) {
	// success
	st := New(1, 2, 3)
	assert.Equal(t, []cell.Cell{1, 2, 3}, st.Cells)
}

func TestEmpty(t *testing.T) {
	// success - true
	b := New().Empty()
	assert.True(t, b)

	// success - false
	b = New(1).Empty()
	assert.False(t, b)
}

func TestLen(t *testing.T) {
	// success
	i := New(1, 2, 3).Len()
	assert.Equal(t, 3, i)
}

func TestPop(t *testing.T) {
	// setup
	st := New(1, 2, 3)

	// success
	c, err := st.Pop()
	assert.Equal(t, cell.Cell(3), c)
	assert.Equal(t, []cell.Cell{1, 2}, st.Cells)
	assert.NoError(t, err)

	// setup
	st = New()

	// error - is empty
	c, err = st.Pop()
	assert.Zero(t, c)
	assert.EqualError(t, err, `cannot pop stack - is empty`)
}

func TestPopSlice(t *testing.T) {
	// setup
	st := New(1, 2, 3)

	// success
	cs, err := st.PopSlice(2)
	assert.Equal(t, []cell.Cell{2, 3}, cs)
	assert.Equal(t, []cell.Cell{1}, st.Cells)
	assert.NoError(t, err)

	// setup
	st = New()

	// error - is empty
	cs, err = st.PopSlice(2)
	assert.Zero(t, cs)
	assert.EqualError(t, err, `cannot pop stack - is insufficient`)
}

func TestPush(t *testing.T) {
	// setup
	st := New(1)

	// success
	st.Push(2)
	assert.Equal(t, []cell.Cell{1, 2}, st.Cells)
}

func TestPushSlice(t *testing.T) {
	// setup
	st := New(1)

	// success
	st.PushSlice([]cell.Cell{2, 3})
	assert.Equal(t, []cell.Cell{1, 2, 3}, st.Cells)
}

func TestString(t *testing.T) {
	// success
	s := New(1, 2, 3).String()
	assert.Equal(t, "1 2 3", s)
}
