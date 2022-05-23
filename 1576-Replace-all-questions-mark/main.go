// space complexity == O(max(n, m)) where n = input size and m = 3
// time complexity == O(n * m) where n = input size and m = 3

package main

import (
	"fmt"
	"strings"
)

func main() {

	input := "a?a"
	modifiedString := modifyString(input)

	fmt.Println("result:", modifiedString)
}

func modifyString(s string) string {

	if len(s) == 0 {
		return ""
	}

	// you only need 3 chars as there is only 3 types of condition that can happen
	poolChar := []string{"a", "b", "c"}
	inputArray := strings.Split(s, "")

	// this test for edge case if there is a single question mark
	if len(inputArray) == 1 && inputArray[0] == "?" {
		return "a"
	}

	for index, value := range inputArray {
		if value == "?" {
			for _, char := range poolChar {
				// this means we are checking for the case where index is on the first value; we only have right neighbor to compare with
				if index == 0 {
					if inputArray[index+1] == char {
						continue
					} else {
						inputArray[index] = char
						break
					}
					// this means we are checking for the case where index is on the last char, where it only has the left value to compare with
				} else if index == len(inputArray)-1 {
					if inputArray[index-1] == char {
						continue
					} else {
						inputArray[index] = char
						break
					}
				} else {
					// this is also the last condition. If the index is not on the first index or the last, it definitely will have both it's left and right neighbors to compare with

					// condition if array @ i-1 and i+1 is the same
					if inputArray[index-1] == inputArray[index+1] {

						// comparison can also be done with inputinputArray[index+1]
						if inputArray[index-1] == char {
							continue
						} else {
							inputArray[index] = char
							break
						}

						// condition if either i-1 or i+1 == char
					} else if inputArray[index-1] == char || inputArray[index+1] == char {
						continue
						// if neither of the conditions apply, then we are certain the char can be assigned to replace the array's ?
					} else {
						inputArray[index] = char
						break
					}
				}
			}
		}
	}
	return strings.Join(inputArray, "")
}
