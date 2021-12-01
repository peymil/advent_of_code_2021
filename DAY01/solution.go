package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var text []int
	for scanner.Scan() {
		scnTxt := scanner.Text()
		value, _ := strconv.ParseInt(scnTxt, 10, 64)
		valueint := int(value)
		text = append(text, valueint)

	}

	fmt.Printf("Answer of the first question : %d\n", findIncreasedValues(text))
	fmt.Printf("Answer of the second question : %d\n", findIncreasedValuesInThreeDimensionalWindow(text))
	file.Close()

}
func findIncreasedValues(arr []int) int {
	// Defining prev without value, resulting the prev being equal to 0.
	var prev int
	var increasedCount int
	for _, value := range arr {
		if value > prev {
			increasedCount++
		}
		prev = value
	}
	// No prev value in first loop. To get the correct result subtract one from result.
	return increasedCount - 1
}
func findIncreasedValuesInThreeDimensionalWindow(arr []int) int {
	//completely wrong
	prevWindowSum := 0
	currWindowSum := 0
	increasedCount := 0
	for _, value := range arr {
		currWindowSum += value
		if currWindowSum > prevWindowSum {
			increasedCount++
		}
		prevWindowSum = currWindowSum
	}

	return increasedCount - 3
}
