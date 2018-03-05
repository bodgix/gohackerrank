package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readData(r io.Reader, sumsCh chan<- int) {
	defer close(sumsCh)

	var ns [3]int
	_, err := fmt.Fscanf(os.Stdin, "%d %d %d", &ns[0], &ns[1], &ns[2])
	if err != nil {
		log.Fatal("Error reading input ", err)
	}

	for _, n := range ns {
		stack, err := readStack(r, n)
		sendSums(stack, sumsCh)
		if err != nil {
			log.Fatal(err)
		}
	}
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

func sendSums(stack []int, sumsCh chan<- int) {
	sum := 0
	for i := len(stack) - 1; i >= 0; i-- {
		sumsCh <- sum + stack[i]
		sum += stack[i]
	}
}

func findMaxSum(sumsCh <-chan int, resultCh chan<- int) {
	defer close(resultCh)
	sums := make(map[int]int)
	maxSum := 0
	for sum := range sumsCh {
		sums[sum]++
		if sums[sum] == 3 {
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	resultCh <- maxSum
}

func main() {
	sums := make(chan int)
	result := make(chan int)

	go readData(os.Stdin, sums)
	go findMaxSum(sums, result)

	fmt.Println(<-result)
}
