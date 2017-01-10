package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

type inputData struct {
	n       int
	queries []int64
}

func readData(in io.Reader) (*inputData, error) {
	var data inputData
	m, err := fmt.Fscanf(in, "%d\n", &data.n)
	if m != 1 {
		return nil, err
	}
	data.queries = make([]int64, 0, data.n)
	data.queries, err = readN(in, data.queries, data.n)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func readN(in io.Reader, result []int64, n int) ([]int64, error) {
	format := "%d\n"
	var j int64
	if n == 0 {
		return result, nil
	} else if n == 1 {
		format = "%d"
	}
	m, err := fmt.Fscanf(in, format, &j)
	if m != 1 {
		return nil, err
	}
	return readN(in, append(result, j), n-1)
}

func bitsNotSet(num int64) []int {
	binary := strconv.FormatInt(num, 2)
	return bitsNotSetRecursively(binary, 0, make([]int, 0, len(binary)))
}

func bitsNotSetRecursively(binary string, i int, result []int) []int {
	if i == len(binary) {
		return result
	}
	if binary[i] == '0' {
		return bitsNotSetRecursively(binary, i+1, append(result, len(binary)-i-1))
	}
	return bitsNotSetRecursively(binary, i+1, result)
}

func numberOfResults(bits []int) int {
	resultsNo := 0
	for _, bitPosition := range bits {
		resultsNo += int(math.Pow(2, float64(bitPosition)))
	}
	return resultsNo
}

func main() {
	data, err := readData(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, number := range data.queries {
		resultsNo := numberOfResults(bitsNotSet(number))
		fmt.Println(resultsNo)
	}
}
