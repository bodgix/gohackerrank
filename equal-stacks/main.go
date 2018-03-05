package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readData(r io.Reader) [3][]int {
	var ns [3]int
	var stack []int
	var sums [3][]int
	_, err := fmt.Fscanf(os.Stdin, "%d %d %d", &ns[0], &ns[1], &ns[2])
	if err != nil {
		log.Fatal("Error reading input ", err)
	}

	for i, n := range ns {
		stack, err = readStack(r, n)
		sums[i] = buildSums(stack)
		if err != nil {
			log.Fatal(err)
		}
	}
	return sums
}

func readStack(r io.Reader, n int) ([]int, error) {
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

func buildSums(stack []int) []int {
	sums := make([]int, 0, len(stack))
	sum := 0
	for i := len(stack) - 1; i >= 0; i-- {
		sums = append(sums, sum+stack[i])
		sum = sum + stack[i]
	}
	return sums
}

func checkCurSum(sum, max int, maxes map[int]int) (map[int]int, int) {
	maxes[sum]++
	if maxes[sum] == 3 {
		if sum > max {
			return maxes, sum
		}
	}
	return maxes, max
}

func findLargestSum(c1, c2, c3 <-chan int, out chan<- int) {
	defer close(out)
	maxes := make(map[int]int)
	max := 0
	for {
		select {
		case sum, ok := <-c1:
			maxes, max = checkCurSum(sum, max, maxes)
			if !ok {
				c1 = nil
			}
		case sum, ok := <-c2:
			maxes, max = checkCurSum(sum, max, maxes)
			if !ok {
				c2 = nil
			}
		case sum, ok := <-c3:
			maxes, max = checkCurSum(sum, max, maxes)
			if !ok {
				c3 = nil
			}
		}
		if c1 == nil && c2 == nil && c3 == nil {
			break
		}
	}
	out <- max
}

func sendStackToChan(c chan<- int, stack []int) {
	defer close(c)
	for _, sum := range stack {
		c <- sum
	}
}

func main() {
	stacks := readData(os.Stdin)
	var chans [3]chan int
	max := make(chan int)

	for i, stack := range stacks {
		chans[i] = make(chan int)
		go sendStackToChan(chans[i], stack)
	}

	go findLargestSum(chans[0], chans[1], chans[2], max)

	result := <-max
	fmt.Println(result)
}
