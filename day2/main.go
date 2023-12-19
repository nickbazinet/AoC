package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data")	
	if err != nil {
		log.Fatalf("Error while reading data file: %+v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sumPart1 int
	var sumPart2 int
	maxAvailableColors := map[string]int{
		"blue": 14,
		"red": 12,
		"green": 13,
	}

	for scanner.Scan() {
		line := scanner.Text()
		game := CreateGame(line)

		if game.IsConfigurationValid(maxAvailableColors) {
			sumPart1 += game.Id
		}

		sumPart2 += game.GetPowerOfSet()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading the data file: %+v", err)
	} else {
		log.Printf("Part 1 - Sum of all the data for PART ONE is: %d", sumPart1)
		log.Printf("Part 2 - Sum of all the power of sets for minimum configuration: %d", sumPart2) 
	}
}
