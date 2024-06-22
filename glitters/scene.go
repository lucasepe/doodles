package glitters

import (
	"context"
	"math"

	"github.com/lucasepe/doodlekit"
)

func Scene() doodlekit.Scene {
	return &scene{
		total: 30,
	}
}

type scene struct {
	total  int
	colors []int
	time   float64
}

func (s *scene) Init(ctx context.Context) {
	s.colors = []int{0, 0, 1, 1, 2, 1, 5, 6, 2, 4, 9, 3, 1, 1, 8, 10}
	s.time = 0
}

func (s *scene) Update(ctx context.Context, _ float64) {
	s.time += 2
}

func (s *scene) Draw(ctx context.Context) {
	rng := doodlekit.Rng(ctx)
	gc := doodlekit.Canvas(ctx)

	for i := 0; i < 400; i++ {
		x, y := rng.RndI(0, 160), rng.RndI(0, 160)

		c := gc.At(x, y)
		if c == 0 {
			c = int(float64(y)/60+math.Abs(math.Mod(float64(x), 32)-16)/60+s.time) % 2

			k := 0.2
			if c == 0 {
				k = 0.8
			}

			if rng.Rnd(0, 1) < k {
				c = 3
			} else {
				c = 11
			}

		} else if rng.Rnd(0, 2) < 1 {
			c = s.colors[c]
		}

		gc.Color(c)
		gc.Circ(int(x), int(y), 1)
	}
}
