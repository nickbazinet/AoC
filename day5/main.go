package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type conversionMap struct {
	name string
	mapping map[int]int
}

func main()  {
	file, err := os.Open("input")	
	if err != nil {
		log.Fatalf("Error while reading data file: %+v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currentName := ""
	currentMapping := make(map[int]int)
	var seeds []string
	var conversionMaps []conversionMap
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "seeds:") {
			for i, seed := range strings.Split(line, " ") {
				if i == 0 {
					continue
				} else {
					seeds = append(seeds, seed)
				}
			}
		} else if strings.HasSuffix(strings.TrimSpace(line), "map:") && currentName == "" {
			currentName = strings.Split(line, " ")[0]
		} else if len(strings.TrimSpace(line)) == 0 {
			if len(currentName) == 0 {
				continue
			}
			conversionMaps = append(conversionMaps, conversionMap{
				name: currentName,
				mapping: currentMapping,
			})
			currentName = ""
			currentMapping = make(map[int]int)
		} else {
			numbers := strings.Split(line, " ")
			rangeLength, _ := strconv.Atoi(numbers[2])
			for i := 0; i < rangeLength; i++ {
				sourceNumber, _ := strconv.Atoi(numbers[1])
				destinationNumber, _ := strconv.Atoi(numbers[0])
				currentMapping[sourceNumber+i] = destinationNumber+i
			}
		}

	}
	conversionMaps = append(conversionMaps, conversionMap{
		name: currentName,
		mapping: currentMapping,
	})
	//log.Println(conversionMaps)

	lowestLocationNumber := -1
	for _, seed := range seeds {
		currentNumber, _ := strconv.Atoi(seed)
		for _, convMap := range conversionMaps {
			if _, ok := convMap.mapping[currentNumber]; ok {
				currentNumber = convMap.mapping[currentNumber]
			} 
		}
		if lowestLocationNumber == -1 || lowestLocationNumber > currentNumber {
			lowestLocationNumber = currentNumber
		}
		//log.Printf("Seed %s correspond to location number %d", seed, currentNumber)
	}
	log.Printf("Part1: Lowest Location number: %d", lowestLocationNumber)
}
