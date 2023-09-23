package rle_encoder

import (
	"strconv"
	"strings"
	"unicode"
)

func Encode(input string) string {
	var str strings.Builder
	l := len(input)

	for i := 0; i < l; i++ {
		count := 1
		for i < l-1 && input[i] == input[i+1] {
			count++
			i++
		}

		str.WriteByte(input[i])
		str.WriteString(strconv.Itoa(count))
	}

	return str.String()
}

func getNumber(input string, index int) (number, ii int) {
	var i int
	for i = index; i < len(input); i++ {
		if !unicode.IsDigit(rune(input[i])) {
			break
		}
	}

	if index == i {
		return 0, 0
	} else {
		str, _ := strconv.Atoi(input[index:i])
		return str, i
	}
}

func Decode(input string) string {
	var str strings.Builder
	l := len(input)

	for i := 0; i < l; {
		c := rune(input[i])
		if !unicode.IsDigit(c) {
			num, index := getNumber(input, i+1)
			str.WriteString(strings.Repeat(string(c), num))
			i = index
		} else {
			i++
		}
	}

	return str.String()
}
