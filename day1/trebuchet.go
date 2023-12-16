package main

import "strconv"
import "log"

// Return the calibration value from a single line. 
// Will return -1 with an error if something goes wrong
func CalculateCalibrationValue(line string) (int, error) {
	var first, last string
	for _, byte := range line {
		str := string(byte)
		_, err := strconv.Atoi(str)
		
		if err == nil {
			if first == "" {
				first = str
				last = str
			} else {
				last = str
			}
		}
	}

	value, err := strconv.Atoi(first + last)
	if err != nil {
		log.Printf("Not able to convert calibration code from string to int: %+v", err)
		return -1, err
	}
	return value, nil
}
