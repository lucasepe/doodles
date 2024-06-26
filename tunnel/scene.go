package tunnel

import (
	"context"
	"math"

	"github.com/lucasepe/doodlekit"
)

func Scene() doodlekit.Scene {
	return &scene{}
}

type scene struct {
	t      float64
	w, h   int
	speed  float64
	colors []int
}

func (s *scene) Init(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	s.speed = 2
	s.colors = []int{4, 5, 2}
	s.w = gc.Width()
	s.h = gc.Height()
}

func (s *scene) Update(ctx context.Context, _ float64) {
	s.t += s.speed
}

func (s *scene) Draw(ctx context.Context) {
	halfW, halfH := s.w/2, s.h/2

	gc := doodlekit.Canvas(ctx)
	gc.Translate(float64(halfW), float64(halfH))

	for y := -halfH; y <= halfH; y++ {
		for x := -halfW; x <= halfW; x++ {
			a := float64(s.t)/16 + (math.Atan2(float64(y), float64(x))+math.Pi)*2.546
			d := 222 / math.Sqrt(float64(x*x+y*y))
			q := s.colors[int(a*12/16)%len(s.colors)]
			var c int
			if d > 1.3*float64(s.h) {
				c = 1
			} else {
				c = q ^ int((d+float64(s.t)/4))%3
			}

			gc.Color(c)
			gc.Pix(x, y)
		}
	}

	gc.Identity()
}
