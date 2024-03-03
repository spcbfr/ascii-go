package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image, err := os.Open("./assets/yusuf.png")
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

	f, err := os.Create("./ascii-result")
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

	os.Exit(0)
}
