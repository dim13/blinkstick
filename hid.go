package blinkstick

import (
	"errors"
	"io"

	"github.com/karalabe/hid"
)

const (
	vendorID  = 0x20a0
	productID = 0x41e5
)

// Errors
var (
	ErrUnsupported = errors.New("unsupproted platform")
	ErrNotFound    = errors.New("device not found")
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
	if !hid.Supported() {
		return nil, ErrUnsupported
	}
	for _, dev := range hid.Enumerate(vendorID, productID) {
		return dev.Open()
	}
	return nil, ErrNotFound
}
