package main

import (
	"math"
)

/*

Dear Reader,

There are 10s of algorithms for calculating the brightness
(sometimes called luminance) of an RGB pixel.

I implement some of them here. feel free to un-comment
any one of them to manually compare their differences

*/

func brightness(pixel Pixel) float64 {

	// Algorithm #1
	// this one simply returns the average of the 3 RGB channels
	// return float64(pixel.R+pixel.G+pixel.B) / 3

	// Algorithm #2
	// Source: https://stackoverflow.com/a/56678483/22768246

	// sRGBtoLin takes a value between 0.0 and 1.0
	// so we map our RGB values to that range
	vR := float64(pixel.R) / 255
	vG := float64(pixel.G) / 255
	vB := float64(pixel.B) / 255

	// luminance usually denoted as Y
	luminance := (0.2126*sRGBtoLin(vR) + 0.7152*sRGBtoLin(vG) + 0.0722*sRGBtoLin(vB))

	// L* is the percieved lightness.
	// the returned value is between 0 -> 100. the rest of the code expects the value
	// to be between 0 -> 255 which I why I convert it below
	Lstar := YtoLstar(luminance) * 255 / 100

	return Lstar

}

func sRGBtoLin(colorChannel float64) float64 {

	// Send this function a decimal sRGB gamma encoded color value
	// between 0.0 and 1.0, and it returns a linearized value.

	if colorChannel <= 0.04045 {
		return colorChannel / 12.92
	} else {
		return math.Pow(((colorChannel + 0.055) / 1.055), 2.4)
	}
}

func YtoLstar(Y float64) float64 {

	// Send this function a luminance value between 0.0 and 1.0,
	// and it returns L* which is "perceptual lightness"

	if Y <= (216.0 / 24389.0) { // The CIE standard states 0.008856 but 216/24389 is the intent for 0.008856451679036
		return Y * (24389 / 27) // The CIE standard states 903.3, but 24389/27 is the intent, making 903.296296296296296
	} else {
		return math.Pow(Y, (1.0/3.0))*116 - 16
	}
}
