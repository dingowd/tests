package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEven(t *testing.T) {
	t.Run("testing of parity", func(t *testing.T) {
		ok := isEven(5)
		require.False(t, ok)
	})
}