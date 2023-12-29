package blinkstick

import (
	"io"

	"github.com/sstallion/go-hid"
)

const (
	vendorID  = 0x20a0
	productID = 0x41e5
)

/* Found device:
   Manufacturer:  Agile Innovative Ltd
   Description:   BlinkStick
   Serial:        BS019296-3.0
   Current Color: #000000
   Mode:          2
   LEDs:          unsupported
   Info Block 1:
   Info Block 2:
*/

// Open blinkstick device
func Open() (io.WriteCloser, error) {
	// Initialize the hid package.
	if err := hid.Init(); err != nil {
		return nil, err
	}
	return hid.OpenFirst(vendorID, productID)
}
