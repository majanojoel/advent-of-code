package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTotalDistance_Example_ShouldMatchExpected(t *testing.T) {
	exampleTestDataPath := "testdata/example1.txt"
	list1, list2, err := parseLists(exampleTestDataPath)
	assert.NoError(t, err)
	expectedResult := 11
	result, err := calculateTotalDistance(list1, list2)
	assert.NoError(t, err)
	assert.Equal(t, result, expectedResult)
}

func TestCalculateSimilarityScore_Example_ShouldMatchExpected(t *testing.T) {
	exampleTestDataPath := "testdata/example1.txt"
	list1, list2, err := parseLists(exampleTestDataPath)
	assert.NoError(t, err)
	expectedResult := 31
	result, err := calculateSimilarityScore(list1, list2)
	assert.NoError(t, err)
	assert.Equal(t, result, expectedResult)
}
