package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateConstant(t *testing.T) {
	testCases := []struct {
		caseName string
		input    int
		expected int
	}{
		{
			caseName: "input less than 100",
			input:    98,
			expected: 0,
		},
		{
			caseName: "input greater than 9998",
			input:    98,
			expected: 0,
		},
		{
			caseName: "should retain 0",
			input:    2111,
			expected: 6174,
		},
		{
			caseName: "should process 3 digit number",
			input:    126,
			expected: 495,
		},
	}

	for _, tCase := range testCases {
		actual := generateConstant(tCase.input)
		assert.Equal(t, tCase.expected, actual)
	}
}

func TestInt2Array(t *testing.T) {
	testCases := []struct {
		caseName string
		input    int
		expected []int
	}{
		{
			caseName: "convert int 2 array",
			input:    12345,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			caseName: "convert int 2 array",
			input:    321,
			expected: []int{3, 2, 1},
		},
	}

	for _, tCase := range testCases {
		actual := int2array(tCase.input)
		assert.Equal(t, tCase.expected, actual)
	}
}

func TestArray2Int(t *testing.T) {
	testCases := []struct {
		caseName string
		input    []int
		expected int
	}{
		{
			caseName: "convert int 2 array",
			input:    []int{1, 2, 3, 4, 5},
			expected: 12345,
		},
		{
			caseName: "convert int 2 array",
			input:    []int{3, 2, 1},
			expected: 321,
		},
	}

	for _, tCase := range testCases {
		actual := array2Int(tCase.input)
		assert.Equal(t, tCase.expected, actual)
	}
}
