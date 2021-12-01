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
	fmt.Printf("Answer of the second question : %d\n", findIncreasedValuesInNWindow(text, 3))
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
func findIncreasedValuesInNWindow(arr []int, window int) int {
	count := 0
	for i := window; i <= len(arr); i++ {
		currSum := sumSlice(arr[i-window : i])
		nextSum := sumSlice(arr[i+1-window : i+1])
		// fmt.Printf("%d,%d", currSum, nextSum)
		if nextSum > currSum {
			count++
		}
	}
	return count
}
func sumSlice(array []int) int {
	var result int
	for _, n := range array {
		result = result + n
	}
	return result
}
