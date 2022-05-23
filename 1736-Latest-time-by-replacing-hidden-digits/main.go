// space complexity O(n) as we make new array
// time complexity O(1) as we are just comparing and assigning logic
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	input := "00:01"
	// we know that we always have length 5
	result := maximumTime(input)
	fmt.Println("this is the result", result)
}

func maximumTime(s string) string {

	if len(s) == 0 || len(s) != 5 {
		return ""
	}

	// make inputArray so it's easier to retrieve data and not easy to work with runes. more complicated
	inputArray := strings.Split(s, "")

	// dealing with the hour section
	if inputArray[0] == "?" && inputArray[1] == "?" {
		inputArray[0] = "2"
		inputArray[1] = "3"
	} else if inputArray[0] == "?" {
		otherValue, _ := strconv.Atoi(inputArray[1])
		if otherValue <= 3 {
			inputArray[0] = "2"
		} else {
			inputArray[0] = "1"
		}
	} else if inputArray[1] == "?" {
		otherValue, _ := strconv.Atoi(inputArray[0])
		if otherValue == 2 {
			inputArray[1] = "3"
		} else {
			inputArray[1] = "9"
		}
	}

	// dealing with the minute section
	if inputArray[3] == "?" && inputArray[4] == "?" {
		inputArray[3] = "5"
		inputArray[4] = "9"
	} else if inputArray[3] == "?" {
		inputArray[3] = "5"
	} else if inputArray[4] == "?" {
		inputArray[4] = "9"
	}

	return strings.Join(inputArray, "")
}
