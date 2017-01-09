package main

import (
	"strings"
	"testing"
)

func samedata(expected, real *inputData) bool {
	if expected.c == real.c && expected.n == real.n && expected.p == real.p {
		for i, v := range real.trips {
			if expected.trips[i] != v {
				return false
			}
		}
		return true
	}
	return false
}

func TestReadDataGood(t *testing.T) {
	inputStr := `5 2 2
1 2 1 4 3`
	expectedOutput := &inputData{5, 2, 2, []int{1, 2, 1, 4, 3}}
	output, err := readData(strings.NewReader(inputStr))
	if err != nil {
		t.Errorf("Reading data returned an error: %v", err)
		t.FailNow()
	}
	if !samedata(expectedOutput, output) {
		t.Errorf("Expected: %v, Got: %v", expectedOutput, output)
	}
}

func TestReadDataBad(t *testing.T) {
	inputStr := `5 2 2 2
1 2 3 4`
	_, err := readData(strings.NewReader(inputStr))
	if err == nil {
		t.Errorf("Expected an error for bad data: %v", inputStr)
	}

	inputStr = `ala ma kota
1 2 3`
	_, err = readData(strings.NewReader(inputStr))
	if err == nil {
		t.Errorf("Expected an error for bad data: %v", inputStr)
	}

	inputStr = `3 1 1
1 2`
	_, err = readData(strings.NewReader(inputStr))
	if err == nil {
		t.Errorf("Expected an error for bad data: %v", inputStr)
	}
}

func TestEnoughBoats(t *testing.T) {
	inputData := &inputData{4, 1, 2, []int{2, 2, 2, 2}}
	if !enoughBoats(inputData) {
		t.Errorf("Expected true for %v", inputData)
	}

	inputData.p = 1
	if enoughBoats(inputData) {
		t.Errorf("Expected false for %v", inputData)
	}
}
