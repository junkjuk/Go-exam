package parser

import (
	"exam/tree"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParser_count(t *testing.T) {
	type testCase struct {
		name           string
		input          string
		expectedResult int
		hasError       bool
	}

	testTable := []testCase{
		{
			name:           "test n-node parsing with count comparing",
			input:          "{\"Root\":{\"Data\":4,\"Children\":[{\"Data\":2,\"Children\":[{\"Data\":3,\"Children\":null}]},{\"Data\":1,\"Children\":null}]}}",
			expectedResult: 4,
			hasError:       false,
		},
		{
			name:           "test b-node parsing with count comparing",
			input:          "{\"Root\":{\"Data\":1,\"Left\":{\"Data\":2,\"Left\":null,\"Right\":null},\"Right\":{\"Data\":3,\"Left\":{\"Data\":4,\"Left\":null,\"Right\":null},\"Right\":{\"Data\":5,\"Left\":null,\"Right\":null}}}}",
			expectedResult: 5,
			hasError:       false,
		},
		{
			name:           "test parsing error string",
			input:          "",
			expectedResult: 0,
			hasError:       true,
		},
		{
			name:           "test parsing error string",
			input:          "{\"Root\":{\"Data\":4,\"Children\":[{\"Data\":2,\"Chi",
			expectedResult: 0,
			hasError:       true,
		},
		{
			name:           "test parsing error string",
			input:          "{\"Root\":{\"Data\":1,\"Left\":{\"Data\":2,\"Left\":null,\"",
			expectedResult: 0,
			hasError:       true,
		},
	}

	for _, test := range testTable {
		actual, err := GetTree(strings.NewReader(test.input))
		if test.hasError {
			assert.NotNil(t, err, test.name)
		} else {
			count := 0
			actual.Traverse(func(node tree.INode) {
				count++
			})
			assert.Equal(t, test.expectedResult, count, test.name)
			assert.Nil(t, err, test.name)
		}
	}
}
