package main

import (
	"log"
	"strconv"
	"strings"
)

// Representation of a single Game of Color Cube
type Game struct {
	Id int
	HighestColorSets map[string]int
}

// Create a game from an entry with the following syntax as an example:
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
func CreateGame(gameEntry string) Game {
	
	highestColorSets := map[string]int{}
	firstSplit := strings.Split(gameEntry, ":")
	gameId, err := strconv.Atoi(string(strings.Split(firstSplit[0], " ")[1]))
	if err != nil {
		log.Fatalf("Not able to convert Game ID from string to int: %+v", err)
	}
	cubeSets := strings.TrimSpace(firstSplit[1])

	for _, set := range strings.Split(cubeSets, ";") {
		cubes := strings.Split(set, ",")

		for _, cube := range cubes {
			
			cubeColorRevealed := strings.Split(strings.TrimSpace(cube), " ")
			numberOfCube, err := strconv.Atoi(cubeColorRevealed[0])
			if err != nil {
				log.Fatalf("Not able to convert Number Of Cube from string to int: %+v", err)
			}

			colorOfCube := cubeColorRevealed[1]

			highest, exists := highestColorSets[colorOfCube]

			if !exists || highest < numberOfCube {
				highestColorSets[colorOfCube] = numberOfCube
			}
		}
	}

	return Game{
		Id: gameId,
		HighestColorSets: highestColorSets,
	}
}

// Return the power of its highest color cube drafted for this current game.
func (g *Game) GetPowerOfSet() int {
	powerOfSet := 1
	for _, highest := range g.HighestColorSets {
		powerOfSet = powerOfSet * highest
	}
	return powerOfSet
}

// Validate if the current game is compatible with the given map of color mapped to 
// its maximum number of color cube associated with.
func (g *Game) IsConfigurationValid(highestAvailableSet map[string]int) bool {
	isValid := true
	for color, maximum := range highestAvailableSet {
		if maximum < g.HighestColorSets[color] {
			isValid = false
		}
	}
	return isValid
}
