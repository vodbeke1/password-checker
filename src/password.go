package main

import (
	"fmt"
	"strings"
)

func stringInSlice(a string, array []string) bool {

	for _, b := range array {
		if b == a {
			return true
		}
	}
	return false
}

func sum(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}
	return sum
}

func passwordCheck(input string) []string {
	password := strings.Split(input, ",")

	var posPasswords []string

	for _, p := range password {
		if 6 <= len(p) && len(p) <= 12 {
			accepted := true
			trigger := make([]int, 3)

			for _, letter := range p {
				var numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
				var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
					"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
					"u", "v", "w", "x", "y", "z"}
				var other = []string{"$", "#", "@"}
				condition := 0

				if stringInSlice(strings.ToLower(string(letter)), alphabet) {
					condition += 1
					trigger[0] = 1
				}
				if stringInSlice(strings.ToLower(string(letter)), numbers) {
					condition += 1
					trigger[1] = 1
				}
				if stringInSlice(strings.ToLower(string(letter)), other) {
					condition += 1
					trigger[2] = 1
				}
				if condition == 0 {
					accepted = false
					break
				}
			}
			if accepted && sum(trigger) == 3 {
				posPasswords = append(posPasswords, p)
			}
		}
	}
	return posPasswords

}

func main() {

	test := "ABd1234@1,a F1#,2w3E*,2We3345"

	fmt.Println(passwordCheck(test))

}
