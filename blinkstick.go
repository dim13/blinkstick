package blinkstick

import (
	"bytes"
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

type Scene [8]color.Color

func rgb(c color.Color) (uint8, uint8, uint8) {
	if c == nil {
		return 0, 0, 0
	}
	r, g, b, _ := c.RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)
}

func SetIndex(w io.Writer, i int, c color.Color) {
	buf := bytes.NewBuffer([]byte{5, 0, byte(i)})
	r, g, b := rgb(c)
	buf.Write([]byte{r, g, b})
	io.Copy(w, buf)
}

func Set(w io.Writer, s Scene) {
	buf := bytes.NewBuffer([]byte{6, 0})
	for _, c := range s {
		r, g, b := rgb(c)
		buf.Write([]byte{g, r, b})
	}
	io.Copy(w, buf)
}

func SetAll(w io.Writer, c color.Color) {
	Set(w, Scene{c, c, c, c, c, c, c, c})
}

func Off(w io.Writer) {
	Set(w, Scene{})
}
