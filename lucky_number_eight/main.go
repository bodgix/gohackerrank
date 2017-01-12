package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

// Helper functions
type inputData struct {
	n, number int
}

func readData(in io.Reader) (*inputData, error) {
	var data inputData
	m, err := fmt.Fscanf(in, "%d\n", &data.n)
	if m != 1 {
		return nil, err
	}
	m, err = fmt.Fscanf(in, "%d", &data.number)
	if m != 1 {
		return nil, err
	}
	return &data, nil
}

func permutations(number string, mask int, result map[int64]bool) (map[int64]bool, error) {
	if float64(mask) == math.Pow(2, float64(len(number))) { // all possibilities checked
		return result, nil
	}
	format := "%" + fmt.Sprintf("0%d", len(number)) + "b"
	maskBase2 := fmt.Sprintf(format, mask)

	var permutation string
	buf := bytes.NewBufferString(permutation)

	for index, bit := range maskBase2 {
		if bit == '1' {
			buf.WriteByte(number[index])
		}
	}
	permutationAsNumber, err := strconv.ParseInt(buf.String(), 10, 64)
	if err != nil {
		return result, err
	}
	result[permutationAsNumber] = true
	return permutations(number, mask+1, result)
}

func howManyDivisable(numbers map[int64]bool, divisor int64) int {
	var result int
	for num := range numbers {
		if math.Remainder(float64(num), float64(divisor)) == 0 {
			result++
		}
	}
	return result
}

// Main function
func main() {
	data, err := readData(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	allSubNumbers := make(map[int64]bool)

	allSubNumbers, err = permutations(fmt.Sprintf("%d", data.number), 1, allSubNumbers)
	if err != nil {
		log.Fatal(err)
	}
	divisorsNo := int64(math.Remainder(float64(howManyDivisable(allSubNumbers, 8)), math.Pow(10.0, 9.0)+7))
	fmt.Println(divisorsNo)
}
