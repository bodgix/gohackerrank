package main

import (
	"strconv"
	"strings"
	"testing"
)

func compareInputData(real, expected *inputData) bool {
	if real.n != expected.n {
		return false
	}
	for i, num := range real.queries {
		if expected.queries[i] != num {
			return false
		}
	}
	return true
}

func TestReadData(t *testing.T) {
	testInput := `3
100
200
300`
	expectedOutput := &inputData{3, []int64{100, 200, 300}}
	output, err := readData(strings.NewReader(testInput))
	if err != nil {
		t.Errorf("readData returned an error: %v", err)
		t.FailNow()
	}
	if !compareInputData(expectedOutput, output) {
		t.Errorf("Expected: %v, got: %v", expectedOutput, output)
	}

	testInput = `3
100
200
300
`
	expectedOutput = &inputData{3, []int64{100, 200, 300}}
	output, err = readData(strings.NewReader(testInput))
	if err != nil {
		t.Errorf("readData returned an error: %v", err)
		t.FailNow()
	}
	if !compareInputData(expectedOutput, output) {
		t.Errorf("Expected: %v, got: %v", expectedOutput, output)
	}
}

func TestReadDataBadData(t *testing.T) {
	testInput := `ala
100
200
300`
	_, err := readData(strings.NewReader(testInput))
	if err == nil {
		t.Error("Expected an error but wasn't returned")
	}
	testInput = `3
ala
ma
kota`
	_, err = readData(strings.NewReader(testInput))
	if err == nil {
		t.Error("Expected an error but wasn't returned")
	}

	testInput = `0
1
2
3`
	output, err := readData(strings.NewReader(testInput))
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
		t.FailNow()
	}
	if output.n != 0 {
		t.Errorf("Expected N to be 0 but got: %v", output.n)
	}
}

func compareIntSlices(expected, got []int) bool {
	if len(expected) != len(got) {
		return false
	}
	for i, num := range expected {
		if got[i] != num {
			return false
		}
	}
	return true
}

func TestFindUnsetBits(t *testing.T) {
	number, _ := strconv.ParseInt("1010", 2, 64)
	got := bitsNotSet(number)
	expected := []int{2, 0}
	if !compareIntSlices(expected, got) {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

	number, _ = strconv.ParseInt("0011010100", 2, 64)
	got = bitsNotSet(number)
	expected = []int{5, 3, 1, 0}
	if !compareIntSlices(expected, got) {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

func TestNumberOfResults(t *testing.T) {
	number := int64(10)
	got := numberOfResults(bitsNotSet(number))
	expected := 5
	if got != expected {
		t.Errorf("Expected %d results for %d, got %d", expected, number, got)
	}
}
