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
	expected := &inputData{3, 968}
	got, err := readData(strings.NewReader(testInput))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		t.FailNow()
	}
	if !compareOutputs(expected, got) {
		t.Errorf("Expected: %v, got %v", expected, got)
	}
}

func TestReadDataBadInput(t *testing.T) {
	testInput := `ala
968`
	_, err := readData(strings.NewReader(testInput))
	if err == nil {
		t.Error("Expected an error but no returned.")
	}

	testInput = `3
ala`
	_, err = readData(strings.NewReader(testInput))
	if err == nil {
		t.Error("Expected an error but no returned.")
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
	got, err := permutations(input, 1, make(map[int64]bool))
	expected := map[int64]bool{8: true, 6: true, 68: true, 9: true, 98: true, 96: true, 968: true}
	if err != nil {
		t.Errorf("Got an error: %v", err)
		t.FailNow()
	}
	if !compareInt64Maps(got, expected) {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

func TestPermutationsBadInput(t *testing.T) {
	input := "test"
	_, err := permutations(input, 1, make(map[int64]bool))
	if err == nil {
		t.Error("Expected an error but wasn't returned")
	}
}

func TestHowManyDivisable(t *testing.T) {
	nums := map[int64]bool{1: true, 3: true, 4: true, 6: true, 7: true, 8: true, 9: true}
	divisable := howManyDivisable(nums, 3)
	if divisable != 3 {
		t.Errorf("Expected 3, got %d", divisable)
	}
}
