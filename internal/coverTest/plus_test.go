package covertest_test

import (
	covertest "main/internal/coverTest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlus(t *testing.T) {
	require.Equal(t, 2, covertest.Plus(1, 1))
}

// func TestMoins(t *testing.T) {
// 	require.Equal(t, 2, covertest.Moins(1, 1))
// }
