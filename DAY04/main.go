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
	var bingoPapers [][]int
	var currentBingoPaper []int
	scanner.Scan()
	bingosStr := strings.Split(scanner.Text(), ",")
	var bingos []int
	for _, bingoStr := range bingosStr {
		bingo64, _ := strconv.ParseInt(bingoStr, 10, 64)
		bingo := int(bingo64)
		bingos = append(bingos, bingo)
	}
	scanner.Scan()
	for scanner.Scan() {
		text := scanner.Text()

		if strings.TrimSpace(text) == "" {
			bingoPapers = append(bingoPapers, currentBingoPaper)
			currentBingoPaper = []int{}
		}
		var word []rune
		for _, char := range text {
			if char != ' ' {
				num, _ := strconv.ParseInt(string(word), 10, 64)
				intnum := int(num)
				currentBingoPaper = append(currentBingoPaper, intnum)
			} else {
				word = append(word, char)
			}
		}
	}
	var bingoCounts []int

	for _, bingoPaper := range bingoPapers {
		bingoCount := 0
		fmt.Printf("%d\n", len(bingoPaper))

		for _, bingoPaperNum := range bingoPaper {
			for _, bingo := range bingos {
				if bingoPaperNum == bingo {
					bingoCount++
				}
			}
		}
		if bingoCount == 50 {
			fmt.Printf("%v\n", bingoCounts)
		}
	}

}
funcIsVerticalOrHorizontal(cubicArray [][]int){
	horizontal
}
