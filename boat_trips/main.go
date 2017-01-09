package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type inputData struct {
	n, c, p int
	trips   []int
}

func readData(input io.Reader) (*inputData, error) {
	data := new(inputData)
	if m, err := fmt.Fscanf(input, "%d %d %d\n", &data.n, &data.c, &data.p); m != 3 {
		return nil, fmt.Errorf("Error reading n, c, p: %v", err)
	}
	var trips []int
	trips, err := readN(input, data.n, data.trips)
	if err != nil {
		return nil, err
	}
	data.trips = trips
	return data, nil
}

func readN(input io.Reader, n int, result []int) ([]int, error) {
	var j int
	if n == 0 {
		return result, nil
	}
	if m, err := fmt.Fscanf(input, "%d", &j); m != 1 {
		return make([]int, 0, 0), err
	}
	return readN(input, n-1, append(result, j))
}

func enoughBoats(data *inputData) bool {
	sort.Ints(data.trips)
	return data.trips[len(data.trips)-1] <= data.c*data.p
}

func main() {
	data, err := readData(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if enoughBoats(data) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
