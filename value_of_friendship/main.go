package main

import (
	"fmt"
	"io"
)

const (
	errUnexpectedNewLine = "unexpected newline"
	errEOF               = "EOF"
)

type friendship struct {
	x, y int
}

type query struct {
	n, m        int
	friendships []friendship
}

type inputData struct {
	q       int
	queries []query
}

func readData(in io.Reader) (*inputData, error) {
	var data inputData
	var line []int
	line, err := readLineOfInts(in)
	if err != nil {
		return nil, err
	}
	data.q = line[0]
	data.queries = make([]query, 0, data.q)

	for qi := data.q; qi > 0; qi-- {
		var curQ query
		nm, err := readLineOfInts(in)
		if err != nil {
			return nil, err
		}
		curQ.n = nm[0]
		curQ.m = nm[1]
		curQ.friendships = make([]friendship, 0, curQ.m)

		for i := curQ.m; i > 0; i-- {
			friends, err := readLineOfInts(in)
			if err != nil {
				return nil, err
			}
			f := friendship{friends[0], friends[1]}
			curQ.friendships = append(curQ.friendships, f)
		}
	}
	return &data, nil
}

func readLineOfInts(in io.Reader) ([]int, error) {
	var d int
	var result []int
	for {
		m, err := fmt.Fscanf(in, "%d", &d)
		if m != 1 {
			if err.Error() == errUnexpectedNewLine || err.Error() == errEOF {
				break
			} else {
				return nil, err
			}
		}
		result = append(result, d)
	}
	return result, nil
}
