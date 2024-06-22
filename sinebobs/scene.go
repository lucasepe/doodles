package sinebobs

import (
	"context"
	"math"

	"github.com/lucasepe/doodlekit"
)

func Scene() doodlekit.Scene {
	return &scene{
		speed: 4.5,
	}
}

type scene struct {
	time  float64
	speed float64
}

func (s *scene) Init(ctx context.Context) {}

func (s *scene) Update(ctx context.Context, dt float64) {
	s.time += s.speed * dt
}

func (s *scene) Draw(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	for l := -7.0; l < 8; l++ {
		x := int(l*11) + 80

		for n := -5.0; n < 6; n++ {
			y := int(n*11) + 80
			p := math.Sin(s.time + l/2 - n/1.5*math.Sin(s.time+n/8))

			gc.Color(2 + int(s.time))
			gc.CircFill(x, y, 2*int(p+2))

			gc.Color(8 + int(s.time))
			gc.CircFill(x, y, 1+int(p))
		}
	}
}
