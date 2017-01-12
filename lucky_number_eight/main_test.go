package main

import (
	"strings"
	"testing"
)

func compareOutputs(expected, got *inputData) bool {
	if expected.n != got.n {
		return false
	}
	if expected.number != got.number {
		return false
	}
	return true
}

func TestReadData(t *testing.T) {
	testInput := `3
968`
	expected := &inputData{3, "968"}
	got, err := readData(strings.NewReader(testInput))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		t.FailNow()
	}
	if !compareOutputs(expected, got) {
		t.Errorf("Expected: %v, got %v", expected, got)
	}
}

func compareInt64Slices(got, expected []int64) bool {
	if len(got) != len(expected) {
		return false
	}
	for i, n := range expected {
		if got[i] != n {
			return false
		}
	}
	return true
}

func compareStringSlices(got, expected []string) bool {
	if len(got) != len(expected) {
		return false
	}
	for i, n := range expected {
		if got[i] != n {
			return false
		}
	}
	return true
}

func compareInt64Maps(got, expected map[int64]bool) bool {
	if len(got) != len(expected) {
		return false
	}
	for n := range expected {
		if !got[n] {
			return false
		}
	}
	return true
}

func TestPermutations(t *testing.T) {
	input := "968"
	got, err := permutations(input, 1, make([]string, 0, 8))
	if err != nil {
		t.Errorf("Got an error: %v", err)
		t.FailNow()
	}
	expected := []string{"8", "6", "68", "9", "98", "96", "968"}
	if !compareStringSlices(got, expected) {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

func TestHowManyDivisable(t *testing.T) {
	nums := []string{"1", "3", "6", "9", "11", "12"}
	divisable, _ := howManyDivisable(nums, 3)
	if divisable != 4 {
		t.Errorf("Expected 3, got %d", divisable)
	}
}

func TestHowManyDivisableBadData(t *testing.T) {
	nums := []string{"1", "3", "ouch!", "9", "11", "12"}
	_, err := howManyDivisable(nums, 3)
	if err == nil {
		t.Error("Expected an error but wasn't returned.")
	}
}
