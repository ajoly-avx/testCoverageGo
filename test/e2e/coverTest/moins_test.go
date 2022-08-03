package covertest_test

import (
	covertest "main/internal/coverTest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoins(t *testing.T) {
	require.Equal(t, 0, covertest.Moins(1, 1))
}
