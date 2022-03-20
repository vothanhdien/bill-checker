package activity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckerStub_Check(t *testing.T) {
	cs := CheckerStub{
		M: make(map[string]int),
	}

	check, _ := cs.Check("FOO_2")
	require.False(t, check)
	check, _ = cs.Check("FOO_2")
	require.False(t, check)
	check, _ = cs.Check("FOO_2")
	require.True(t, check)

	check, _ = cs.Check("FOO")
	require.True(t, check)
	check, _ = cs.Check("FOO_1")
	require.True(t, check)
}
