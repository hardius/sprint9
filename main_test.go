package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateFunc(t *testing.T) {
	genArr := generateRandomElements(0)
	require.Nil(t, genArr)

	sizes := []int{1, 100, 1984, 3213, 777_777, 100_000_000}
	for i := range sizes {
		genArr := generateRandomElements(sizes[i])
		require.NotNil(t, genArr)
	}
}

func TestMaximum(t *testing.T) {
	genArr := [][]int{generateRandomElements(0), generateRandomElements(1)}
	cases := []struct {
		expected int
		actual   int
	}{
		{0, maximum(genArr[0])},
		{genArr[1][0], maximum(genArr[1])},
	}

	for _, c := range cases {
		require.Equal(t, c.expected, c.actual)
	}
}
func TestMaximumAndMaxChunks(t *testing.T) {
	sizes := []int{0, 1, 100, 1984, 3213, 777_777, 100_000_000}

	for i := range sizes {
		genArr := generateRandomElements(sizes[i])
		require.Equal(t, maximum(genArr), maxChunks(genArr))
	}
}
