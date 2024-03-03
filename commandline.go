package main

import (
	"flag"
	"fmt"
)

const usage = `Usage of ascii-go:
	-i, --input <file>   The image to be converted
	-o, --output <file>  The result file
	-h, --help    prints this message

	note: if no output file is given, the program prints to stdout
`

func cmdFlags() (string, string) {

	var input, output string
	flag.StringVar(&input, "input", "", "The image to be converted")
	flag.StringVar(&output, "output", "", "The result file")

	flag.StringVar(&input, "i", "", "The image to be converted")
	flag.StringVar(&output, "o", "", "The result file")

	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()

	return input, output

}
