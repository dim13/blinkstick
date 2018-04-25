package main

import (
	"image/color"
	"log"
	"time"

	"dim13.org/blinkstick"
)

func newBinary(n int, c color.Color) blinkstick.Frame {
	var s blinkstick.Frame
	for i := 0; i < 8; i++ {
		if n&(1<<uint(i)) != 0 {
			s[i] = c
		}
	}
	return s
}

func newTwiddle(n int, c color.Color) blinkstick.Frame {
	var s blinkstick.Frame
	n %= 14
	if n < 8 {
		s[n] = c
	} else {
		s[14-n] = c
	}
	return s
}

func main() {
	dev, err := blinkstick.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer dev.Close()

	defer blinkstick.Off(dev)

	pal := []color.Color{
		color.YCbCr{0x1f, 0x00, 0xff}, // red
		color.YCbCr{0x3f, 0x00, 0xbf}, // yellow
		color.YCbCr{0x1f, 0x00, 0x00}, // green
		color.YCbCr{0x1f, 0xff, 0x1f}, // blue
		color.Black,                   // off
	}

	for _, c := range pal {
		for i := 0; i < 8; i++ {
			blinkstick.SetIndex(dev, i, c)
			time.Sleep(time.Second / 2)
		}
	}
}
