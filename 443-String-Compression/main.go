// time complexity O(n) and space complexity O(1) since we don't extra space
package main

import (
	"fmt"
	"strconv"
)

func main() {

	input := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}

	compressedLength := compress(input)

	fmt.Println("the compressed length", compressedLength)
	fmt.Println(string(input))

}

func compress(chars []byte) int {

	var stringConverted string
	index := 0
	i := 0

	for i < len(chars) {

		j := i // always restart j to be the same as i

		for j < len(chars) && chars[j] == chars[i] {
			j++
		}

		if j-i > 1 {
			stringConverted = convertInttoString(j - i)
			for _, v := range stringConverted {
				index++
				chars[index] = byte(v)
			}
		}

		i = j
		index++
	}
	return index
}

// helper function not even needed since the build in library can do the same
func convertInttoString(i int) string {
	var stringConverted string

	stringConverted = strconv.Itoa(i)

	return stringConverted
}
