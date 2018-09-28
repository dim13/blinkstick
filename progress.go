package blinkstick

import (
	"image/color"
	"io"
	"time"
)

var (
	red    = color.YCbCr{0x1f, 0x00, 0xff}
	yellow = color.YCbCr{0x3f, 0x00, 0xbf}
	green  = color.YCbCr{0x1f, 0x00, 0x00}
	blue   = color.YCbCr{0x1f, 0xff, 0x1f}
	black  = color.Black
)

type Progress struct {
	//f     Frame
	start time.Time
	soft  time.Duration
	hard  time.Duration
	end   time.Duration
	n     int
}

func NewProgress(soft, hard, end time.Duration) Progress {
	return Progress{
		start: time.Now(),
		soft:  soft,
		hard:  hard,
		end:   end,
	}
}

func (p *Progress) Update(w io.Writer) {
	done := time.Since(p.start)
	switch {
	case p.hard < done:
		//p.f[p.n] = red
	case p.soft < done:
		//p.f[p.n] = yellow
	default:
		//p.f[p.n] = green
	}
	//Set(w, p.f)
	p.n = (p.n + 1) % 8
}
