package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMulOps_SimpleTest_ShouldReturnOne(t *testing.T) {
	expected := mulOp{
		X: 654,
		Y: 789,
	}
	actual := parseMemoryForMulOps("mul(654,789)")
	assert.Len(t, actual, 1)
	assert.Equal(t, expected.X, actual[0].X)
	assert.Equal(t, expected.Y, actual[0].Y)
}

func TestCalculateSum_Example_ShouldReturnExpected(t *testing.T) {
	testFilePath := "testdata/example1.txt"
	ops, err := parseMultiplicationOperations(testFilePath)
	assert.NoError(t, err)
	assert.Len(t, ops, 4)
	actualSum := getSumFromMulOps(ops)
	assert.Equal(t, 161, actualSum)
}
