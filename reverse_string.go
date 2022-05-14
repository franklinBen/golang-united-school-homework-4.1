package reverse_string

import "fmt"

func ReverseString(input string) (output string) {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	for i := 0; i < len(runes); i++ {
		output = fmt.Sprint(runes[i])
	}
	return output
}