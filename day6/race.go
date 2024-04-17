package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Race struct to hold the time and distance
type Race struct {
	Time     int
	Distance int
}

func readRacesFromFile(filename string) ([]Race, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	races := []Race{}

	// Assuming the first line is "Time" and the second line is "Distance"
	timeLine, distanceLine := "", ""
	if scanner.Scan() {
		timeLine = scanner.Text()
	}
	if scanner.Scan() {
		distanceLine = scanner.Text()
	}

	// Removing the prefix to get the raw data
	timeData := strings.TrimPrefix(timeLine, "Time:")
	distanceData := strings.TrimPrefix(distanceLine, "Distance:")

	timeValues := strings.Fields(timeData)
	distanceValues := strings.Fields(distanceData)

	// Assuming both lines have the same number of values
	for i := 0; i < len(timeValues); i++ {
		time, _ := strconv.Atoi(timeValues[i])
		distance, _ := strconv.Atoi(distanceValues[i])

		race := Race{
			Time:     time,
			Distance: distance,
		}

		races = append(races, race)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return races, nil
}
