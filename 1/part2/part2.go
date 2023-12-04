package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func numStringToInt(str string) []int32 {
	// I'm sure there's a better way to do this
	var output []int32
	for i := 0; i < len(str); i++ {
		if i <= len(str)-5 {
			var substr string = str[i:]
			switch {
			case substr[0:5] == "three":
				output = append(output, 3)
			case substr[0:5] == "seven":
				output = append(output, 7)
			case substr[0:5] == "eight":
				output = append(output, 8)
			}
		}
		if i <= len(str)-4 {
			var substr string = str[i:]
			switch {
			case substr[0:4] == "four":
				output = append(output, 4)
			case substr[0:4] == "five":
				output = append(output, 5)
			case substr[0:4] == "nine":
				output = append(output, 9)
			}
		}
		if i <= len(str)-3 {
			var substr string = str[i:]
			switch {
			case substr[0:3] == "one":
				output = append(output, 1)
			case substr[0:3] == "two":
				output = append(output, 2)
			case substr[0:3] == "six":
				output = append(output, 6)
			}
		}
		// Mutually exclusive with any written int
		if (48 <= str[i]) && (str[i] <= 57) {
			output = append(output, int32(str[i])-48)
		}
	}
	return output
}

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
		ints := numStringToInt(line)
		if len(ints) == 1 {
			num := ints[0]*10 + ints[0]
			calibrationValues = append(calibrationValues, num)
		} else if len(ints) >= 2 {
			num := ints[0]*10 + ints[len(ints)-1]
			calibrationValues = append(calibrationValues, num)
		} else {
			log.Fatal("Could not parse any integers from line")
		}
	}
	for i := 0; i < len(calibrationValues); i++ {
		sum += calibrationValues[i]
	}
	fmt.Println(sum)
}
