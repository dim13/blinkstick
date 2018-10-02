// Package blinckstick implements BlinkStick Stripe HID interface
package blinkstick

import (
	"errors"
	"image/color"
	"io"
)

/* Reports:
   1: LED Data [R, G, B]
   2: Name [Binary Data 0..32]
   3: Data [Binary Data 0..32]
   4: Mode set [MODE]: 0 - RGB LED Strip, 1 - Inverse RGB LED Strip, 2 - WS2812
   5: LED Data [CHANNEL, INDEX, R, G, B]
   6: LED Frame [Channel, [G, R, B][0..7]]
   7: LED Frame [Channel, [G, R, B][0..15]]
   8: LED Frame [Channel, [G, R, B][0..31]]
   9: LED Frame [Channel, [G, R, B][0..63]]
*/

func rgb(c color.Color) (r, g, b uint8) {
	if c == nil {
		return 0, 0, 0
	}
	R, G, B, _ := c.RGBA()
	return uint8(R >> 8), uint8(G >> 8), uint8(B >> 8)
}

// SetIndex sets color by index
func SetIndex(w io.Writer, i int, c color.Color) error {
	r, g, b := rgb(c)
	_, err := w.Write([]byte{5, 0, uint8(i), r, g, b})
	return err
}

// Set 0 to 64 colors
func Set(w io.Writer, colors ...color.Color) error {
	var buf []byte
	switch l := len(colors); {
	case l <= 8:
		buf = make([]byte, 3*8+2)
		buf[0] = 6
	case l <= 16:
		buf = make([]byte, 3*16+2)
		buf[0] = 7
	case l <= 32:
		buf = make([]byte, 3*32+2)
		buf[0] = 8
	case l <= 64:
		buf = make([]byte, 3*64+2)
		buf[0] = 9
	default:
		return errors.New("too many colors")
	}
	for i, c := range colors {
		r, g, b := rgb(c)
		buf[3*i+2] = g
		buf[3*i+3] = r
		buf[3*i+4] = b
	}
	_, err := w.Write(buf)
	return err
}

// SetAll sets all n LEDs to same color
func SetAll(w io.Writer, n int, c color.Color) error {
	if n == 0 {
		n = 8
	}
	colors := make([]color.Color, n)
	for i := 0; i < n; i++ {
		colors[i] = c
	}
	return Set(w, colors...)
}

// Off all n LEDs
func Off(w io.Writer, n int) error {
	return SetAll(w, n, nil)
}
