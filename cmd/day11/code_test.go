package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	data, err := os.ReadFile("../../examples/11a.txt")
	assert.NoError(t, err)
	if assert.NotEmpty(t, data) {
		result, ok := partOne(data)
		assert.True(t, ok)
		assert.Equal(t, uint64(5), result)
	}
}

func TestPartTwo(t *testing.T) {
	data, err := os.ReadFile("../../examples/11b.txt")
	assert.NoError(t, err)
	if assert.NotEmpty(t, data) {
		result, ok := partTwo(data)
		assert.True(t, ok)
		assert.Equal(t, uint64(2), result)
	}
}
