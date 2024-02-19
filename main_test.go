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
