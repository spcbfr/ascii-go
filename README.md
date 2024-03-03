I started learning golang last week, and this is my first project using it.

the program takes a PNG image and converts it to ascii art.

## How does it work?

First I split the image into pixels, each pixel has an RGB value. I use the RGB value to determine the brightness of a each pixel. finally I map the brightness to a character from an ascii character set and output the result to a file


