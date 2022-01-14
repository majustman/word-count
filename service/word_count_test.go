package service

import (
	"reflect"
	"testing"
)

type testDefinition struct {
	input string
	expectedResult map[string]int
}

func TestWordCount(t *testing.T) {
	tests := []testDefinition{
		{
			input: "Golang Maps is a collection of unordered pairs of key-value",
			expectedResult: map[string]int{"Golang": 1, "Maps": 1, "is": 1, "a": 1, "collection": 1,
				"of": 2, "unordered": 1, "pairs": 1, "key-value": 1},
		},
		{
			input: "Golang Golang Golang     Maps is a collection of unordered pairs of key-value",
			expectedResult: map[string]int{"Golang": 3, "Maps": 1, "is": 1, "a": 1, "collection": 1,
				"of": 2, "unordered": 1, "pairs": 1, "key-value": 1},
		},
		{
			input: "\n\n\n\nGolang \n Maps \t is\ta\ncollection\tof unordered\npairs of key-value",
			expectedResult: map[string]int{"Golang": 1, "Maps": 1, "is": 1, "a": 1, "collection": 1,
				"of": 2, "unordered": 1, "pairs": 1, "key-value": 1},
		},
		{
			input: "Golang Maps - - - - - a collection of unordered pairs of key-value",
			expectedResult: map[string]int{"Golang": 1, "Maps": 1, "a": 1, "collection": 1,
				"of": 2, "unordered": 1, "pairs": 1, "key-value": 1, "-": 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if !reflect.DeepEqual(tt.expectedResult, WordCount(tt.input)) {
				t.Fatalf("test for %v is failed", tt.input)
			}
		})
	}
}
