package glitters

import (
	"context"
	"math"

	"github.com/lucasepe/doodlekit"
)

func Scene(total int) doodlekit.Scene {
	return &scene{
		total: total,
	}
}

type scene struct {
	w, h   int
	total  int
	colors []int
	time   float64
}

func (s *scene) Init(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	s.w, s.h = gc.Width(), gc.Height()
	if s.total < 0 {
		s.total = 40
	}

	if s.total > 2*s.w {
		s.total = s.w
	}

	s.colors = []int{1, 1, 2, 2, 4, 2, 6, 7, 4, 5, 10, 4, 2, 2, 9, 11}
	s.time = 0
}

func (s *scene) Update(ctx context.Context, _ float64) {
	s.time += 2
}

func (s *scene) Draw(ctx context.Context) {
	rng := doodlekit.Rng(ctx)
	gc := doodlekit.Canvas(ctx)

	for i := 0; i < s.total; i++ {
		x, y := rng.RndI(0, s.w), rng.RndI(0, s.h)

		c := gc.At(x, y)
		if c <= 1 {
			c = int(float64(y)/60+math.Abs(math.Mod(float64(x), 32)-16)/60+s.time) % 2

			k := 0.2
			if c == 0 {
				k = 0.8
			}

			if rng.Rnd(0, 1) < k {
				c = 4
			} else {
				c = 12
			}

		} else if rng.Rnd(0, 2) < 1 {
			c = s.colors[c]
		}

		gc.Color(c)
		gc.Circ(int(x), int(y), 1)
	}
}
