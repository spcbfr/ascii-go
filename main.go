package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {

	inputFile, outputFile := cmdFlags()

	if inputFile == "" {
		fmt.Println("An input file is required")
		fmt.Print(usage)
	}

	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	image, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer image.Close()

	pixels, err := getPixels(image)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ascii := getAscii(pixels)

	if outputFile != "" {
		f, err := os.Create(outputFile)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer f.Close()
		_, err = f.WriteString(ascii)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println(ascii)
	}

	os.Exit(0)
}
