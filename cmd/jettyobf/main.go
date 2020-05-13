package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ecnepsnai/jettyobf"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage '%s' <Obfuscated password OR plain-text password>\n", os.Args[0])
		os.Exit(1)
	}
	input := os.Args[1]

	if strings.HasPrefix(input, "OBF:") {
		fmt.Println(jettyobf.Deobfuscate(input))
	} else {
		fmt.Println(jettyobf.Obfuscate(input))
	}
}
