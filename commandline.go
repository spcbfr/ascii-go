package main

import (
	"flag"
	"fmt"
)

const FLAG_USAGE = `Usage of ascii-go:
	-i, --input <file>   The image to be converted
	-o, --output <file>  The result file
	-c, --invert inverts the image colors
	-h, --help    prints this message

	note: if no output file is given, the program prints to stdout
`

func cmdFlags() (string, string, bool) {

	var input, output string
	var inverted bool

	flag.StringVar(&input, "input", "", "The image to be converted")
	flag.StringVar(&output, "output", "", "The result file")
	flag.BoolVar(&inverted, "invert", false, "Invert the colors of output")

	flag.StringVar(&input, "i", "", "The image to be converted")
	flag.StringVar(&output, "o", "", "The result file")
	flag.BoolVar(&inverted, "c", false, "Invert the colors of output")

	flag.Usage = func() { fmt.Print(FLAG_USAGE) }
	flag.Parse()

	return input, output, inverted

}
