package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("USAGE: missing string")
		return
	}

	var input string = strings.Join(os.Args[1:], " ")

	hexa := hex.EncodeToString([]byte(input))

	fmt.Println(hexa)
}
