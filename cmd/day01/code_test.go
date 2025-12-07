package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	data, err := os.ReadFile("../../examples/1.txt")
	assert.NoError(t, err)
	if assert.NotEmpty(t, data) {
		result, ok := partOne(data)
		assert.True(t, ok)
		assert.Equal(t, uint64(3), result)
	}
}

func TestPartTwo(t *testing.T) {
	data, err := os.ReadFile("../../examples/1.txt")
	assert.NoError(t, err)
	if assert.NotEmpty(t, data) {
		result, ok := partTwo(data)
		assert.True(t, ok)
		assert.Equal(t, uint64(6), result)
	}
}

func BenchmarkPartOne(b *testing.B) {
	data, err := os.ReadFile("../../examples/1.txt")
	if err != nil {
		fmt.Println("unable to read file:", err.Error())
		os.Exit(1)
	}

	b.ReportAllocs()

	for b.Loop() {
		partOne(data)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	data, err := os.ReadFile("../../examples/1.txt")
	if err != nil {
		fmt.Println("unable to read file:", err.Error())
		os.Exit(1)
	}

	b.ReportAllocs()

	for b.Loop() {
		partTwo(data)
	}
}
