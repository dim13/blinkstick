// Package blinckstick implements BlinkStick Stripe HID interface
package blinkstick

import (
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

// SetIndex sets color by index
func SetIndex(w io.Writer, i int, c color.Color) {
	if c == nil {
		c = color.Black
	}
	r, g, b, _ := c.RGBA()
	w.Write([]byte{5, 0, uint8(i), uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)})
}

func Set(w io.Writer, colors ...color.Color) {
	var buf []byte
	switch {
	case len(colors) <= 8:
		buf = make([]byte, 3*8+2)
		buf[0] = 6
	case len(colors) <= 16:
		buf = make([]byte, 3*16+2)
		buf[0] = 7
	case len(colors) <= 32:
		buf = make([]byte, 3*32+2)
		buf[0] = 8
	case len(colors) <= 64:
		buf = make([]byte, 3*64+2)
		buf[0] = 9
	default:
		panic("too many colors")
	}
	for i, c := range colors {
		if c == nil {
			c = color.Black
		}
		r, g, b, _ := c.RGBA()
		buf[3*i+2] = uint8(g >> 8)
		buf[3*i+3] = uint8(r >> 8)
		buf[3*i+4] = uint8(b >> 8)
	}
	w.Write(buf)
}

// SetAll sets all (8) LEDs to same color
func SetAll(w io.Writer, c color.Color) {
	Set(w, c, c, c, c, c, c, c, c)
}

// Off all LEDs
func Off(w io.Writer) {
	SetAll(w, color.Black)
}
