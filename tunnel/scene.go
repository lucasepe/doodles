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
	s.colors = []int{3, 4, 1}
	s.w = gc.Width() / 2
	s.h = gc.Height() / 2
}

func (s *scene) Update(ctx context.Context, _ float64) {
	s.t += s.speed
}

func (s *scene) Draw(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	for y := -s.h; y <= s.h; y++ {
		for x := -s.w; x <= s.w; x++ {
			a := float64(s.t)/16 + (math.Atan2(float64(y), float64(x))+math.Pi)*2.546
			d := 999 / math.Sqrt(float64(x*x+y*y)+1)
			q := s.colors[1+int(a*12/16)%(len(s.colors)-1)]
			var c int
			if d > float64(s.h) {
				c = 0
			} else {
				c = q ^ int((d+float64(s.t)/4))%3
			}

			gc.Color(c)
			gc.Pix(s.w+x, s.h+y)
		}
	}
}
