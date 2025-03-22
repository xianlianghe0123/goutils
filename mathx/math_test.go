package mathx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAbs(t *testing.T) {
	require.Equal(t, 1, Abs(1))
	require.Equal(t, 0, Abs(0))
	require.Equal(t, 3, Abs(-3))
	require.Equal(t, 2.71, Abs(2.71))
	require.Equal(t, 0., Abs(0.00))
	require.Equal(t, 3.14, Abs(-3.14))
}

func TestSum(t *testing.T) {
	require.Equal(t, 15, Sum(1, 2, 3, 4, 5))
	require.Equal(t, 11., Sum(1.1, 2.2, 3.3, 4.4))
	require.Equal(t, 4-2i, Sum(1+2i, 3-4i))
}
