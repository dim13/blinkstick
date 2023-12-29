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

	for i, c := range palette.Plan9 {
		blinkstick.Set(dev, newTwiddle(i, c)...)
		time.Sleep(time.Second / 20)
	}
}
