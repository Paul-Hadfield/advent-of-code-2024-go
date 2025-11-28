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

func getCounts(values []int) map[int]int {
	counts := make(map[int]int, len(values))
	for _, v := range values {
		counts[v]++
	}
	return counts
}

func totalSimilarities(leftValues []int, rightCounts map[int]int) int {
	total := 0
	for _, v := range leftValues {
		total += v * rightCounts[v]
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
	rightValueCounts := getCounts(rightValues)

	total := totalSimilarities(leftValues, rightValueCounts)

	fmt.Println(total)
	fmt.Printf("run time: %s\n", time.Since(start))
}
