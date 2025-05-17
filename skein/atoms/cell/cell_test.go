package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	c := New(1.23)
	assert.Equal(t, Cell(1.23), c)
}

func TestParse(t *testing.T) {
	// success - integer
	c, err := Parse("1")
	assert.Equal(t, Cell(1), c)
	assert.NoError(t, err)

	// success - float
	c, err = Parse("1.23")
	assert.Equal(t, Cell(1.23), c)
	assert.NoError(t, err)

	// error - invalid format
	c, err = Parse("")
	assert.Zero(t, c)
	assert.EqualError(t, err, `cannot parse cell "" - invalid format`)
}

func TestBool(t *testing.T) {
	// success - true
	b := New(1.23).Bool()
	assert.True(t, b)

	// success - false
	b = New(0).Bool()
	assert.False(t, b)
}

func TestNative(t *testing.T) {
	// success
	f := New(1.23).Native()
	assert.Equal(t, float64(1.23), f)
}

func TestString(t *testing.T) {
	// success - integer
	s := New(1).String()
	assert.Equal(t, "1", s)

	// success - float
	s = New(1.23).String()
	assert.Equal(t, "1.23", s)
}
