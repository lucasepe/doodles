package stars

import (
	"context"

	"github.com/lucasepe/doodlekit"
)

func Scene() doodlekit.Scene {
	return &scene{
		speed: 5,
		total: 100,
	}
}

type scene struct {
	stars  []star
	colors []int
	total  int
	speed  float64
}

func (s *scene) Init(ctx context.Context) {
	rng := doodlekit.Rng(ctx)

	s.colors = []int{6, 12, 1}

	s.stars = make([]star, s.total)
	for i := 0; i < s.total; i++ {
		s.stars[i].x = rng.Rnd(0, 160)
		s.stars[i].y = rng.Rnd(0, 160)
		s.stars[i].vx = rng.Rnd(0.5, 3)
	}
}

func (s *scene) Update(ctx context.Context, dt float64) {
	rng := doodlekit.Rng(ctx)

	for i := 0; i < s.total; i++ {
		s.stars[i].x = s.stars[i].x - s.speed*s.stars[i].vx

		if s.stars[i].x < 0 {
			s.stars[i].x = 160
			s.stars[i].y = rng.Rnd(0, 160)
			s.stars[i].vx = rng.Rnd(0.5, 3)
		}
	}
}

func (s *scene) Draw(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	for _, el := range s.stars {
		gc.Color(s.colors[int(el.vx)])
		gc.Pix(int(el.x), int(el.y))
	}
}

type star struct {
	x, y float64
	vx   float64
}
