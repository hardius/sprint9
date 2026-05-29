package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateFunc(t *testing.T) {
	genArr := generateRandomElements(0)
	require.Equal(t, []int{}, genArr)
}
func TestMaxDefaultCases(t *testing.T) {
	genArr := generateRandomElements(100000)
	require.Equal(t, maximum(genArr), maxChunks(genArr))

	genArr = generateRandomElements(350607)
	require.Equal(t, maximum(genArr), maxChunks(genArr))
}

func TestMaxZeroCase(t *testing.T) {
	genArr := generateRandomElements(0)
	require.Equal(t, 0, maximum(genArr))
	require.Equal(t, maximum(genArr), maxChunks(genArr))
}

func TestMaxSingleCase(t *testing.T) {
	genArr := generateRandomElements(1)
	require.Equal(t, genArr[0], maximum(genArr))
	require.Equal(t, maximum(genArr), maxChunks(genArr))
}
