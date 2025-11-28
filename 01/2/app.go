package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"time"
)

func orderArray(values []int) []int {
	sort.Ints(values)
	return values
}

func totalDistance(leftValues, rightValues []int) int {
	if len(leftValues) != len(rightValues) {
		log.Fatalf("left and right slices differ in length: %d vs %d", len(leftValues), len(rightValues))
	}

	total := 0
	for i := 0; i < len(leftValues); i++ {
		diff := leftValues[i] - rightValues[i]
		if diff < 0 {
			diff = -diff
		}
		total += diff
	}
	return total
}

func main() {
	start := time.Now()

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("failed to open data.txt: %v", err)
	}
	defer file.Close()

	leftValues := make([]int, 0, 1024)
	rightValues := make([]int, 0, 1024)

	reader := bufio.NewReader(file)
	for {
		var leftValue, rightValue int
		_, err := fmt.Fscan(reader, &leftValue, &rightValue)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to scan values: %v", err)
		}
		leftValues = append(leftValues, leftValue)
		rightValues = append(rightValues, rightValue)
	}

	leftValues = orderArray(leftValues)
	rightValues = orderArray(rightValues)

	total := totalDistance(leftValues, rightValues)

	fmt.Println(total)
	fmt.Printf("run time: %s\n", time.Since(start))
}
