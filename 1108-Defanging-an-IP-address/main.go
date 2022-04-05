// time complexity O(n) while space complexity O(n)
// or your argument for both cases can be constant because technically ipv4 and ipv6 address are fixed. former is 32 while latter is 128 bits. In other words they are fixed length!
//11010000 01100010 11000000 10101010 = 208.98.192.170 as you can see the max number of numbers only == 12 + 3 dots = 15 chars to range to!
// if we know that the input is in sets of 3, it's even better. We can immediately go into the index and replace them since we know they are valid ipv4 addresses.

package main

import "strings"

func main() {

	input := "1.1.1.1"
	defangIPaddr(input)

}

func defangIPaddr(address string) string {
	const beforeReplacement = "."
	const afterReplacement = "[.]"

	convertedInput := strings.Split(address, "")

	for k, v := range convertedInput {

		if v == beforeReplacement {
			convertedInput[k] = afterReplacement
		}
	}
	return strings.Join((convertedInput), "")
}
