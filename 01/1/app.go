package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func splitLine(line string) (leftValue, rightValue int) {

	parts := strings.Fields(line)
	if len(parts) < 2 {
		log.Fatalf("line %q does not contain two values", line)
	}

	leftValue, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("invalid left value %q: %v", parts[0], err)
	}

	rightValue, err = strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("invalid right value %q: %v", parts[1], err)
	}

	return leftValue, rightValue
}

func orderArray(values []int) []int {
	sort.Ints(values)
	return values
}

func measureDistances(leftValues, rightValues []int) []int {
	if len(leftValues) != len(rightValues) {
		log.Fatalf("left and right slices differ in length: %d vs %d", len(leftValues), len(rightValues))
	}

	distances := make([]int, 0, len(leftValues))
	for i := 0; i < len(leftValues); i++ {
		if leftValues[i] > rightValues[i] {
			distances = append(distances, leftValues[i]-rightValues[i])
		} else {
			distances = append(distances, rightValues[i]-leftValues[i])
		}
	}

	return distances
}

func sumDistances(distances []int) int {
	total := 0
	for _, d := range distances {
		total += d
	}
	return total
}

func main() {
	start := time.Now()

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("failed to open example.txt: %v", err)
	}
	defer file.Close()

	var leftValues []int
	var rightValues []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		leftValue, rightValue := splitLine(scanner.Text())
		leftValues = append(leftValues, leftValue)
		rightValues = append(rightValues, rightValue)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read example.txt: %v", err)
	}

	leftValues = orderArray(leftValues)
	rightValues = orderArray(rightValues)

	distances := measureDistances(leftValues, rightValues)

	total := sumDistances(distances)

	fmt.Println(total)
	fmt.Printf("run time: %s\n", time.Since(start))
}
