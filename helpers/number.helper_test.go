package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertFloat64ToInt64(t *testing.T) {
	floadnum := float64(110.0000001)

	s := ConvertFloat64ToInt64(floadnum)

	assert.Equal(t, int64(110), s)
}
