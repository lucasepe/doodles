package xmastree

import (
	"context"
	"math"

	"github.com/lucasepe/doodlekit"
)

func Scene() doodlekit.Scene {
	return &scene{
		speed: 3,
		t:     1,
	}
}

type scene struct {
	w, h  int
	t     float64
	speed float64
}

func (s *scene) Init(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	s.w, s.h = gc.Width(), gc.Height()
}

func (s *scene) Update(ctx context.Context, _ float64) {
	s.t += s.speed
}

func (s *scene) Draw(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	for i := 0.0; i < 800; i++ {
		a := math.Pi * 2 / 60 * (i + s.t/3)
		z := math.Cos(a)*i/10 + 500

		x := math.Sin(a)*i/6*float64(s.h)/z + float64(s.w/2)
		y := (i/3)*float64(s.h)/z + float64(s.h)/4
		r := i / 500

		gc.Color(1 + int(i+s.t)>>3)
		gc.CircFill(int(x), int(y), int(r))
	}
}
