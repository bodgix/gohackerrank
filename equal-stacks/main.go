package main

import (
	"fmt"
	"io"
)

func readStack(r io.Reader, n int) ([]int, error) {
	res := make([]int, 0, n)
	sum := 0

	for i := 0; i < n; i++ {
		var b int
		_, err := fmt.Fscanf(r, "%d", &b)
		if err != nil {
			return nil, err
		}
		res = append(res, b+sum)
		sum = b + sum
	}
	return res, nil
}

func findLargestSum(c1, c2, c3 <-chan int, out chan<- int) {
	maxes := make(map[int]int)
	max := 0
	for {
		var sum int
		select {
		case sum, ok := <-c1:
			if !ok {
				c1 = nil
			}
		case sum, ok := <-c2:
			if !ok {
				c2 = nil
			}
		case sum, ok := <-c3:
			if !ok {
				c3 = nil
			}
		}
		maxes[sum]++
		if maxes[sum] == 3 {
			if sum > max {
				max = sum
			}
		}
		if c1 == nil && c2 == nil && c3 == nil {
			break
		}
	}
	out <- max
}

func main() {
	stacks := readData
	var height, highestStack int
	var equal bool
	for equal, height, highestStack = areStacksEqual(stacks); height > 0 && !equal; equal, height, highestStack = areStacksEqual(stacks) {
		stacks[highestStack].pop()
	}
	fmt.Println(height)
}
