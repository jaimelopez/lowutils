package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const H_START_NOTATION_NUM string = "0x"
const H_START_NOTATION_SYM string = "\\x"
const H_START_NOTATION_CHAR string = "x"
const H_END_NOTATION string = "h"

const CHARS_NUM int = 2

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("USAGE: missing hexadecimal")
		return
	}

	var results []string

	for _, value := range os.Args[1:] {
		chunk := sanetize(value)

		bs, err := hex.DecodeString(chunk)

		if err != nil {
			results = append(results, fmt.Sprintf("<Err:%s>", chunk))
			continue
		}

		results = append(results, string(bs))
	}

	fmt.Println(strings.Join(results, " "))
}

func sanetize(input string) string {
	result := ""

	for input != "" {
		if len(input) <= CHARS_NUM {
			result += input
			input = ""

			continue
		}

		prefix := ""

		if strings.HasPrefix(input, H_START_NOTATION_NUM) {
			prefix = H_START_NOTATION_NUM
		} else if strings.HasPrefix(input, H_START_NOTATION_SYM) {
			prefix = H_START_NOTATION_SYM
		} else if strings.HasPrefix(input, H_START_NOTATION_CHAR) {
			prefix = H_START_NOTATION_CHAR
		}

		if prefix != "" {
			result += input[len(prefix) : len(prefix)+CHARS_NUM]
			input = input[len(prefix)+CHARS_NUM:]

			continue
		}

		if input[CHARS_NUM:CHARS_NUM+len(H_END_NOTATION)] == H_END_NOTATION {
			result += input[:CHARS_NUM]
			input = input[CHARS_NUM+len(H_END_NOTATION):]

			continue
		}

		result += input[:CHARS_NUM]
		input = input[CHARS_NUM:]
	}

	return result
}
