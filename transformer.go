package main

import (
	"image"
	"io"
)

func getAsciiChar(brightness int, invert bool) string {
	const codes = "Ñ@#W$9876543210?!abc;:+=-,._          "
	var charCode int

	if invert {
		charCode = int(mapRange(float64(brightness), 0, 255, 0, float64(len(codes)-1)))
	} else {
		charCode = int(mapRange(float64(brightness), 0, 255, float64(len(codes)), 0))
	}

	return string(codes[charCode])
}

func getAscii(data [][]Pixel, invert bool) string {
	var result string
	for y := 0; y < len(data); y++ {
		var line string
		for x := 0; x < len(data[y]); x++ {

			line = line + getAsciiChar(brightness(data[y][x]), invert)
		}
		result = result + line + "\n"
	}
	return result
}

func getPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	var pixels [][]Pixel
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}
	return pixels, nil
}

func rgbaToPixel(r uint32, g uint32, b uint32, _ uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257)}
}
func brightness(pixel Pixel) int {
	// placeholder calculation method
	return (pixel.R + pixel.G + pixel.B) / 3
}

type Pixel struct {
	R int
	G int
	B int
}

func mapRange(value, minInput, maxInput, minOutput, maxOutput float64) float64 {
	// Clamping the value to the input range
	if value < minInput {
		value = minInput
	} else if value > maxInput {
		value = maxInput
	}

	// Calculate the proportion of the value in the input range
	proportion := (value - minInput) / (maxInput - minInput)

	// Map the proportion to the output range
	newValue := proportion*(maxOutput-minOutput) + minOutput

	return newValue
}
