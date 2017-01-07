package main

import (
	"strings"
	"testing"
)

func TestReadData(t *testing.T) {
	// readData parses the input data
	testInput := `4
2 2 3 3`
	testReader := strings.NewReader(testInput)
	testOutput := readData(testReader)
	expectedOutput := socksStock{total: 4}
	expectedOutput.socks = make(map[int]int)
	expectedOutput.socks[2] = 2
	expectedOutput.socks[3] = 2

	if testOutput.total != expectedOutput.total {
		t.Errorf("Expected %d total, got %d", expectedOutput.total, testOutput.total)
	}
	for color, count := range expectedOutput.socks {
		if count != testOutput.socks[color] {
			t.Errorf("Expected %d socks of color %d, got: %d", count, color, testOutput.socks[color])
		}
	}
}

func TestNumberOfPairs(t *testing.T) {
	// number of pairs should return 2
	testStock := socksStock{
		total: 5,
		socks: map[int]int{
			1: 1,
			2: 2,
			3: 2,
		},
	}
	testResult := numberOfPairs(testStock)
	if testResult != 2 {
		t.Errorf("Expected 2 pairs, got: %d", testResult)
	}
}
