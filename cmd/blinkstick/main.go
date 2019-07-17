package main

import (
	"image/color"
	"image/color/palette"
	"log"
	"time"

	"github.com/dim13/blinkstick"
)

func newBinary(n int, c color.Color) []color.Color {
	s := make([]color.Color, 8)
	for i := 0; i < 8; i++ {
		if n&(1<<uint(i)) != 0 {
			s[i] = c
		}
	}
	return s
}

func newTwiddle(n int, c color.Color) []color.Color {
	s := make([]color.Color, 8)
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

	defer blinkstick.Off(dev, 8)

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

	for i := 0; ; i++ {
		c := palette.Plan9[i%256]
		f := newTwiddle(i, c)
		blinkstick.Set(dev, f...)
		time.Sleep(time.Second / 14)
	}
}
