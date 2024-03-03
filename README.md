I started learning golang last week (April 2024), and this is my first project using it.

the program takes a PNG image and converts it to ascii art.

## How does it work?

First I split the image into pixels, each pixel has an RGB value. I use the RGB value to determine the brightness of a each pixel. finally I map the brightness to a character from an ascii character set and output the result to a file

## How to use

1. Install [golang](https://go.dev/doc/install)
2. clone the project `git clone https://github.com/spcbfr/ascii-go.git`
3. run `go build .`
4. run `./ascii-go` (`ascii-go.exe` if you're on windows) you should see the usage instructions

The result is best viewed with a zoomed-out text editor with no-syntax highlighting, a low line-height improves the effect further too.

**Note:** if the output doesn't look totally right, try using the `--invert` flag. this uses denser characters for lower brightness pixels and vice-versa

Happy Ascii-ing!
