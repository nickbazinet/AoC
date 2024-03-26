package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	cardId string
	winningNumber []int
	currentNumber []int
}

func main()  {
	file, err := os.Open("input")	
	if err != nil {
		log.Fatalf("Error while reading data file: %+v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cards := []Card{}
	for scanner.Scan() {
		line := scanner.Text()
		cardId := strings.Split(line, ":")[0]
		numbers := strings.Split(line, ":")[1]
		
		winningNumberStr := strings.TrimSpace(strings.Split(numbers, "|")[0])
		myNumberAsString := strings.TrimSpace(strings.Split(numbers, "|")[1])

		currentCard := Card {
			cardId: cardId,
			winningNumber: convertSliceOfStringToSliceOfInt(strings.Split(winningNumberStr, " ")),
			currentNumber: convertSliceOfStringToSliceOfInt(strings.Split(myNumberAsString, " ")),
		}
		cards = append(cards, currentCard)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading the data file: %+v", err)
	}

	gamePts := float64(0)
	for _, card := range cards {
		numberOfMatches := 0
		for _, number := range card.currentNumber {
			if contains(card.winningNumber, number) && number != 0 {
				numberOfMatches += 1
			}
		}
		if numberOfMatches != 0 {
			gamePts += math.Pow(2, float64(numberOfMatches - 1))
		}
	}
	log.Printf("Number of Points for the who game: %.0f", gamePts)
}

func convertSliceOfStringToSliceOfInt(input []string) []int  {
	ints := make([]int, len(input))
	for i, s := range input {
		ints[i], _ = strconv.Atoi(s)	
	}
	return ints
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}


