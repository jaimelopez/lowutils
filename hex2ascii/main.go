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

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("USAGE: missing hexadecimal")
		return
	}

	var results []string

	for _, value := range os.Args[1:] {
		var input string = sanetize(value)

		bs, err := hex.DecodeString(input)

		if err != nil {
			results = append(results, fmt.Sprintf("<Err:%s>", input))
			continue
		}

		results = append(results, string(bs))
	}

	fmt.Println(strings.Join(results, " "))
}

func sanetize(input string) string {
	if strings.HasPrefix(input, H_START_NOTATION_NUM) {
		return input[len(H_START_NOTATION_NUM):]
	}

	if strings.HasPrefix(input, H_START_NOTATION_SYM) {
		return input[len(H_START_NOTATION_SYM):]
	}

	if strings.HasPrefix(input, H_START_NOTATION_CHAR) {
		return input[len(H_START_NOTATION_CHAR):]
	}

	if strings.HasSuffix(input, H_END_NOTATION) {
		return input[:len(input)-len(H_END_NOTATION)]
	}

	return input
}
