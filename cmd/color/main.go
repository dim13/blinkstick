package main

import (
	"flag"
	"log"

	"dim13.org/blinkstick"
	"golang.org/x/image/colornames"
)

func main() {
	flag.Parse()
	c, ok := colornames.Map[flag.Arg(0)]
	if !ok {
		return
	}

	dev, err := blinkstick.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer dev.Close()

	blinkstick.SetAll(dev, 8, c)
}
