package math

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, int64(2), addInt64(1,1))
}
