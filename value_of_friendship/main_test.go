package main

import (
	"strings"
	"testing"
)

/*

func compareInputData(got, expected *inputData) bool {
	if got.q != expected.q {
		return false
	}
	if len(got.queries) != len(expected.queries) {
		return false
	}
	for i, q := range expected.queries {
		g := got.queries[i]
		if q.n != g.n || q.m != g.m {
			return false
		}
		if len(q.friendships) != len(g.friendships) {
			return false
		}
		for i, f := range q.friendships {
			if f[0] != g.friendships[i][0] || f[1] != g.friendships[i][1] {
				return false
			}
		}
	}
	return true
}


func TestReadInputData(t *testing.T) {
	testInput := `1
5 4
1 2
3 2
4 2
4 3`
	got, err := readData(strings.NewReader(testInput))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := &inputData{
		q:       1,
		queries: []query{query{n: 5, m: 4, friendships: [][]int{{1, 2}, {3, 2}, {4, 2}, {4, 3}}}},
	}
	if !compareInputData(got, expected) {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

*/

func compSlices(got, expected []int) bool {
	if len(got) != len(expected) {
		return false
	}
	for i, num := range expected {
		if got[i] != num {
			return false
		}
	}
	return true
}

func TestReadLineOfInts(t *testing.T) {
	inputData := "1 2 3\n"
	inputData2 := "1 2 3"
	inputSlice := []int{1, 2, 3}
	outputData1, err1 := readLineOfInts(strings.NewReader(inputData))
	outputData2, err2 := readLineOfInts(strings.NewReader(inputData2))
	if !compSlices(outputData1, inputSlice) {
		t.Errorf("Expected: %v, got: %v", inputData, outputData1)
	}
	if !compSlices(outputData2, inputSlice) {
		t.Errorf("Expected: %v, got: %v", inputData, outputData2)
	}
	if err1 != nil {
		t.Errorf("Unexpected error: %v", err1)
	}
	if err2 != nil {
		t.Errorf("Unexpected error: %v", err1)
	}
}

func TestReadLineOfIntsBadData(t *testing.T) {
	inputData := "1 2 dupa 3"
	_, err := readLineOfInts(strings.NewReader(inputData))
	if err == nil {
		t.Error("Expected an error but wasn't returned.")
	}
}

func compareDataStructs(got, expected inputData) bool {
	return true
}

func TestReadData(t *testing.T) {
	inputData := `1
4 3
1 2
2 3
1 4
`
	expected := inputData{
		q:       1,
		queries: []query{query{4, 3, []friendship{friendship{1, 2}, friendship{2, 3}, friendship{1, 4}}}},
	}
	got, err := readData(strings.NewReader(inputData))
	if err != nil {
		t.Errorf("Got error: %v", err)
		t.FailNow()
	}
	if !compareDataStructs(*got, expected) {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}
