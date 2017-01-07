package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type socksStock struct {
	total int
	socks map[int]int
}

func readData(input io.Reader) (*socksStock, error) {
	var stock socksStock
	_, err := fmt.Fscanf(input, "%d\n", &stock.total)
	if err != nil {
		return nil, err
	}
	stock.socks = make(map[int]int)
	if err = readN(input, stock.socks, 0, stock.total); err != nil {
		return nil, err
	}
	return &stock, nil
}

func readN(input io.Reader, result map[int]int, i, n int) error {
	if n == 0 {
		return nil
	}
	var j int
	if m, err := fmt.Fscanf(input, "%d", &j); m != 1 {
		return fmt.Errorf("Error reading %d-th number: %v", i, err)
	}
	result[j]++
	return readN(input, result, i+1, n-1)
}

func numberOfPairs(stock *socksStock) int {
	var pairs int
	for _, count := range stock.socks {
		pairs += count / 2
	}
	return pairs
}

func main() {
	stock, err := readData(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numberOfPairs(stock))
}
