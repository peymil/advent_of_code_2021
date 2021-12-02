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
	depth := 0
	horizontal := 0
	depthWAim := 0
	aim := 0
	for scanner.Scan() {
		text := scanner.Text()
		splittedText := strings.Split(text, " ")
		//
		command := splittedText[0]
		_value, _ := strconv.ParseInt(splittedText[1], 10, 64)
		value := int(_value)
		if command == "forward" {
			horizontal += value
			depthWAim -= aim * value
		}
		if command == "up" {
			depth -= value
			aim += value
		}
		if command == "down" {
			depth += value
			aim -= value
		}

	}
	fmt.Printf("Answer of the first question: %d\n", depth*horizontal)
	fmt.Printf("Answer of the second question: %d\n", depthWAim*horizontal)
}
