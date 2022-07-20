package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestColor(t *testing.T) {
	t.Parallel()
	t.Run("values", func(t *testing.T) {
		t.Parallel()
		require.Len(t, values, int(count))
		// Start at 1, because 0 is black and should be 0x000000 anyways
		for c := Enum(1); c < count; c++ {
			assert.NotZero(t, c.Value(), "color.Enum(%d)", c)
		}
	})
	t.Run("names", func(t *testing.T) {
		t.Parallel()
		require.Len(t, names, int(count))
		for c := Enum(0); c < count; c++ {
			assert.NotZero(t, c.Name(), "color.Enum(%d)", c)
		}
	})
	t.Run("defaults", func(t *testing.T) {
		t.Parallel()
		require.Len(t, defaults, int(count))
		for c := Enum(1); c < count; c++ {
			assert.NotZero(t, defaults[c], "color.Enum(%d)", c)
		}
	})
}
