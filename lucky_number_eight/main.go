package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

// Helper functions
type inputData struct {
	n      int
	number string
}

func readData(in io.Reader) (*inputData, error) {
	var data inputData
	m, err := fmt.Fscanf(in, "%d\n", &data.n)
	if m != 1 {
		return nil, err
	}
	m, err = fmt.Fscanf(in, "%s", &data.number)
	if m != 1 {
		return nil, err
	}
	return &data, nil
}

func permutations(number string, mask int, result []string) ([]string, error) {
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
	return permutations(number, mask+1, append(result, buf.String()))
}

func howManyDivisable(numbers []string, divisor int64) (int, error) {
	var result int
	for _, numStr := range numbers {
		var num int64
		m, err := fmt.Sscanf(numStr, "%d", &num)
		if m != 1 {
			return -1, err
		}
		if math.Remainder(float64(num), float64(divisor)) == 0 {
			result++
		}
	}
	return result, nil
}

// Main function
func main() {
	data, err := readData(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading input data: %v", err)
	}
	digitsNo := len(data.number)
	allSubNumbers := make([]string, 0, int(math.Pow(2.0, float64(digitsNo))))

	allSubNumbers, err = permutations(data.number, 1, allSubNumbers)
	if err != nil {
		log.Fatal(err)
	}

	divisorsNo, err := howManyDivisable(allSubNumbers, 8)
	if err != nil {
		log.Fatal(err)
	}
	divisorsNoModulo := int64(math.Remainder(float64(divisorsNo), math.Pow(10.0, 9.0)+7))
	fmt.Println(divisorsNoModulo)
}
