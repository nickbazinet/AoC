package main

import (
	"log"
	"strconv"
	"strings"
)

var letter_to_digits_map = map[string]string {
	"one": "1",
	"two": "2",
	"three": "3",
	"four": "4",
	"five": "5",
	"six": "6",
	"seven": "7",
	"eight": "8",
	"nine": "9",
}

var starting_letter_digit = []string{"o", "t", "f", "s", "e", "n"}

type CalibrationDigit struct {
	Digit string
	Position int
}

// Return the calibration value from a single line. 
// Will return -1 with an error if something goes wrong
func CalculateCalibrationValueV2(line string) (int, error) {
	var first, last CalibrationDigit
	for i, byte := range line {
		currentLetter := string(byte)
		_, err := strconv.Atoi(currentLetter)
		
		if err == nil {
			updateBothCalibrationValues(&first, &last, currentLetter, i)
		} else {

			// Don't waste your time searching when no digit start with the current letter
			if !isInSlice(currentLetter, starting_letter_digit) {
				continue
			}
			
			textToSearch := line[i:]
			for key, value := range letter_to_digits_map {

				// Don't waste your time to search for a digit that doesn't start with the same letter 
				// as the one we are searching for
				if string(key[0]) != currentLetter {
					continue
				}
				
				// The index of the letter digit to find should be 0 (start of this string)
				if strings.Index(textToSearch, key) != 0 {
					continue
				} else {
					updateBothCalibrationValues(&first, &last, value, i)
				}
			}
		}
	}

	value, err := strconv.Atoi(first.Digit + last.Digit)
	if err != nil {
		log.Printf("Not able to convert calibration code from string to int: %+v", err)
		return -1, err
	}
	return value, nil
}

// Return a boolean to know if the target value existing in the given slice
func isInSlice(target string, slice []string) bool {
    for _, element := range slice {
        if element == target {
            return true
        }
    }
    return false
}

// Update the Digit and Position values with the one givin in parameters
func (c *CalibrationDigit) UpdateCalibrationValues(digit string, position int) {
	c.Digit = digit
	c.Position = position
}

// Update both First and Last CalibrationDigit with the given parameters. 
// If the first CalibrationDigit is already initialized, nothing will be changed on it
func updateBothCalibrationValues(first, last *CalibrationDigit, digit string, position int) {
	if *first == (CalibrationDigit{}) {
		first.UpdateCalibrationValues(digit, position)
	}
	last.UpdateCalibrationValues(digit, position)
}
