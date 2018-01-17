package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const (
	hStartNotationNum  string = "0x"
	hStartNotationSym  string = "\\x"
	hStartNotationChar string = "x"
	hEndNotation       string = "h"

	charsNum int = 2
)

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
		if len(input) <= charsNum {
			result += input
			input = ""

			continue
		}

		prefix := ""

		if strings.HasPrefix(input, hStartNotationNum) {
			prefix = hStartNotationNum
		} else if strings.HasPrefix(input, hStartNotationSym) {
			prefix = hStartNotationSym
		} else if strings.HasPrefix(input, hStartNotationChar) {
			prefix = hStartNotationChar
		}

		if prefix != "" {
			result += input[len(prefix) : len(prefix)+charsNum]
			input = input[len(prefix)+charsNum:]

			continue
		}

		if input[charsNum:charsNum+len(hEndNotation)] == hEndNotation {
			result += input[:charsNum]
			input = input[charsNum+len(hEndNotation):]

			continue
		}

		result += input[:charsNum]
		input = input[charsNum:]
	}

	return result
}
