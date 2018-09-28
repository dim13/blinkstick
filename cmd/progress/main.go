package main

import (
	"log"
	"time"

	"dim13.org/blinkstick"
)

func main() {
	dev, err := blinkstick.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer dev.Close()

	p := NewProgress(5*time.Second, 7*time.Second, 8*time.Second)
	for i := 0; i < 8; i++ {
		blinkstick.Set(dev, p.Update()...)
		time.Sleep(time.Second)
	}

	blinkstick.Off(dev)
}
