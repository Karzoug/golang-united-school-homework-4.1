package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	res := make([]rune, len(runes))

	for i := len(runes) - 1; i >= 0; i-- {
		res = append(res, runes[i])
	}
	output = string(res)
	return output
}
