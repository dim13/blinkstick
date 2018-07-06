// +build ignore

package blinkstick

import (
	"testing"
	"time"
)

func TestProgress(t *testing.T) {
	dev, err := Open()
	if err != nil {
		t.Fatal(err)
	}
	defer dev.Close()

	p := NewProgress(5*time.Second, 7*time.Second, 8*time.Second)
	for i := 0; i < 8; i++ {
		p.Update(dev)
		time.Sleep(time.Second)
	}

	Off(dev)
}
