package main

import (
	"fmt"
	"log"
	"math"
)


func main()  {
	filename := "inputv2"
	races, err := readRacesFromFile(filename)
	if err != nil {
		log.Fatalf("Error reading the file: %+v", err)
	}

	result := 1
	for _, race := range races {
		s1, s2 := solveEquations(float64(race.Time), float64(race.Distance))
		fmt.Printf("Minimum of time holding the button where we can beat the record: %v\n", s2)
		fmt.Printf("Amount of time that the boat is going to sail: %v\n", s1)
		fmt.Println("================================================")

		nextWinningNumber := math.Ceil(s2)
		if nextWinningNumber == s2 {
			nextWinningNumber += 1
		}

		numberOfHigherCombinasion := ((float64(race.Time)/2)-nextWinningNumber)*2
		if race.Time+1%2 != 0 {
			numberOfHigherCombinasion += 1
		}

		fmt.Printf("Number of winning combinaison: %f\n", numberOfHigherCombinasion)
		fmt.Println("================================================")
		result *= int(numberOfHigherCombinasion)
	}
	fmt.Printf("The final Result is: %d\n", result)
}
