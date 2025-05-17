package atom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomise(t *testing.T) {
	// error - invalid format
	a, err := Atomise("")
	assert.Nil(t, a)
	assert.EqualError(t, err, `cannot atomise "" - invalid format`)
}
