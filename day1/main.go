package main

import (
	"bufio"
	"log"
	"os"
)

func main()  {

	file, err := os.Open("data")	
	if err != nil {
		log.Fatalf("Error while reading data file: %+v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int 
	for scanner.Scan() {
		line := scanner.Text()
		value, err := CalculateCalibrationValue(line)
		if err != nil {
			log.Fatalf("Error while reading the single line: %s. Error: %+v", line, err)
		} else {
			sum = value + sum
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading the data file: %+v", err)
	} else {
		log.Printf("Sum of all the data file's line is: %d", sum)
	}
}
