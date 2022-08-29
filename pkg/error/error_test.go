package error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertError(t *testing.T) {
	e, ok := ConvertError(ErrNotFound.String())
	assert.Equal(t, true, ok)
	assert.Equal(t, ErrNotFound, e)

}
