I started learning golang last week, and this is my first project using it.

the program takes a PNG image and converts it to ascii art.

## How does it work?

First I split the image into pixels, each pixel has an RGB value. I use the RGB value to determine the brightness of a each pixel. finally I map the brightness to a character from an ascii character set and output the result to a file

## How to use

1. Install [golang](https://go.dev/doc/install)
2. clone the project `git clone https://github.com/spcbfr/ascii-go.git`
3. run `go build .`
4. run `./ascii-go` or (`ascii-go.exe` if you're on windows) you should see the usage instructions

Happy Ascii-ing!
