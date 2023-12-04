package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// define sum variable for end
	var sum int32
	sum = 0
	//read in file
	calibrationDoc, err := os.Open("../input/input")
	//handle read error
	if err != nil {
		log.Fatal(err)
	}
	//instantiate new scanner
	scanner := bufio.NewScanner(calibrationDoc)
	//split into lines
	scanner.Split(bufio.ScanLines)
	//define new slice to contain calibration values
	var calibrationValues []int32
	//scan through calibrationDoc line by line
	for scanner.Scan() {
		var line = scanner.Text()
		//create new slice to hold line's calibration values
		var ints []int32
		for i := 0; i < len(line); i++ {
			// Check ascii value of every character to check for integers
			if (48 <= line[i]) && (line[i] <= 57) {
				ints = append(ints, int32(line[i])-48) // Convert ascii int to regular int
			} else {
				continue
			}
		}
		if len(ints) == 1 {
			calibrationValues = append(calibrationValues, ints[0]*10+ints[0])
		} else if len(ints) >= 2 {
			calibrationValues = append(calibrationValues, ints[0]*10+ints[len(ints)-1])
		} else {
			log.Fatal("Could not parse any integers from line")
		}
	}
	for i := 0; i < len(calibrationValues); i++ {
		sum += calibrationValues[i]
	}
	fmt.Println(sum)
}
