package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var lines uint16
	var onecounts [12]uint16
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "")
		for n, char := range text {
			if char == "1" {
				onecounts[n]++
			}
		}
		lines++
	}
	var gammaSlice []rune
	var opsilonSlice []rune
	fmt.Printf("%v", onecounts)
	for _, onecount := range onecounts {
		fmt.Printf("%d\n", onecount)
		if onecount >= lines/2 {
			gammaSlice = append(gammaSlice, '1')
			opsilonSlice = append(opsilonSlice, '0')
		} else {
			gammaSlice = append(gammaSlice, '0')
			opsilonSlice = append(opsilonSlice, '1')
		}
	}
	gamma, _ := strconv.ParseInt(string(gammaSlice), 2, 32)
	opsilon, _ := strconv.ParseInt(string(opsilonSlice), 2, 32)
	fmt.Printf("Answer of the first question: %d\n", opsilon*gamma)
}
