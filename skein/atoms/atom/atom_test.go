package atom

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/skein/skein/atoms/cell"
)

func TestAtomise(t *testing.T) {
	// success - cell
	a, err := Atomise("1.23")
	assert.Equal(t, cell.Cell(1.23), a)

	// error - invalid format
	a, err = Atomise("")
	assert.Nil(t, a)
	assert.EqualError(t, err, `cannot atomise "" - invalid format`)
}
