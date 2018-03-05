package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type stackElement struct {
	b      int
	height int
}

type stack []stackElement

func (s stack) peek() stackElement {
	if len(s) == 0 {
		return stackElement{}
	}
	return s[len(s)-1]
}

func (s *stack) pop() stackElement {
	if len(*s) == 0 {
		return stackElement{}
	}
	stackElement := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return stackElement
}

func buildStack(elems []int) stack {
	s := make(stack, 0, len(elems))
	height := 0
	for i := len(elems) - 1; i >= 0; i-- {
		height += elems[i]
		s = append(s, stackElement{elems[i], height})
	}
	return s
}

func readnInt(r io.Reader, n int) ([]int, error) {
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var b int
		_, err := fmt.Fscanf(r, "%d", &b)
		if err != nil {
			return nil, err
		}
		res = append(res, b)
	}
	return res, nil
}

func readData(r io.Reader) [3]stack {
	var ns [3]int
	var stacks [3]stack
	_, err := fmt.Fscanf(os.Stdin, "%d %d %d", &ns[0], &ns[1], &ns[2])
	if err != nil {
		log.Fatal("Error reading input ", err)
	}

	for i, n := range ns {
		input, err := readnInt(r, n)
		if err != nil {
			log.Fatal(err)
		}
		stacks[i] = buildStack(input)
	}
	return stacks
}

func areStacksEqual(stacks [3]stack) (bool, int, int) {
	firstHeight := false
	height := 0
	equal := true
	heighestStack := 0

	for stackNo, s := range stacks {
		if firstHeight {
			if len(s) == 0 {
				height = 0
				heighestStack = stackNo
			} else {
				height = s[len(s)-1].height
				heighestStack = stackNo
			}
		} else {
			var newHeight int
			if len(s) == 0 {
				newHeight = 0
			} else {
				newHeight = s[len(s)-1].height
			}
			if height == newHeight {
				equal = true
			} else if height < newHeight {
				equal = false
				heighestStack = stackNo
				height = newHeight
			} else {
				equal = false
			}
		}
	}
	return equal, height, heighestStack
}

func main() {
	stacks := readData(os.Stdin)
	var height, highestStack int
	var equal bool
	for equal, height, highestStack = areStacksEqual(stacks); height > 0 && !equal; equal, height, highestStack = areStacksEqual(stacks) {
		stacks[highestStack].pop()
	}
	fmt.Println(height)
}
