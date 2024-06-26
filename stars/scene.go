package stars

import (
	"context"

	"github.com/lucasepe/doodlekit"
)

func Scene(total int) doodlekit.Scene {
	return &scene{
		total: total,
	}
}

type scene struct {
	w, h   int
	stars  []star
	colors []int
	total  int
	speed  float64
}

func (s *scene) Init(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)
	rng := doodlekit.Rng(ctx)

	if s.total <= 0 {
		s.total = 100
	}

	s.w, s.h = gc.Width(), gc.Height()

	s.speed = 1.5
	s.colors = []int{7, 13, 2}

	s.stars = make([]star, s.total)
	for i := 0; i < s.total; i++ {
		s.stars[i].x = rng.RndI(0, s.w)
		s.stars[i].y = rng.RndI(0, s.h)
		s.stars[i].vx = rng.Rnd(0.4, 2)
	}
}

func (s *scene) Update(ctx context.Context, dt float64) {
	rng := doodlekit.Rng(ctx)

	for i := 0; i < s.total; i++ {
		s.stars[i].x = s.stars[i].x - int(s.speed*s.stars[i].vx)

		if s.stars[i].x < 0 {
			s.stars[i].x = s.w
			s.stars[i].y = rng.RndI(0, s.h)
			s.stars[i].vx = rng.Rnd(0.4, 2)
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
	x, y int
	vx   float64
}
