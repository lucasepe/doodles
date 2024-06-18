package connections

import (
	"context"
	"math"

	"github.com/lucasepe/doodlekit"
)

func Scene() doodlekit.Scene {
	return &scene{
		w:     160,
		h:     160,
		total: 30,
	}
}

type scene struct {
	w, h  int
	total int
	dots  []dot
}

func (s *scene) Init(ctx context.Context) {
	rng := doodlekit.Rng(ctx)

	s.dots = make([]dot, s.total)
	for i := 0; i < s.total; i++ {
		s.dots[i] = dot{
			x:  rng.RndI(s.w),
			y:  rng.RndI(s.h),
			dx: rng.RndI(2) + 1,
			dy: rng.RndI(2) + 1,
			r:  rng.RndI(2) + 2,
		}
	}
}

func (s *scene) Update(ctx context.Context, dt float64) {
	for i := 0; i < s.total; i++ {
		r := 1 + s.dots[i].r
		if s.dots[i].x < r || s.dots[i].x > (s.w-r) {
			s.dots[i].dx = -s.dots[i].dx
		}

		if s.dots[i].y < r || s.dots[i].y > (s.h-r) {
			s.dots[i].dy = -s.dots[i].dy
		}
	}
}

func (s *scene) Draw(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	for i := 0; i < s.total; i++ {
		for j := 0; j < s.total; j++ {
			dx := float64(s.dots[j].x - s.dots[i].x)
			dy := float64(s.dots[j].y - s.dots[i].y)
			d := math.Hypot(dx, dy)
			dr := float64(s.dots[i].r + s.dots[j].r)
			if d < (2 + 3*dr) {
				gc.Color(11)
				gc.Line(s.dots[j].x, s.dots[j].y, s.dots[i].x, s.dots[i].y)

				if d <= dr {
					s.dots[j].dx, s.dots[i].dx = s.dots[i].dx, s.dots[j].dx
					s.dots[j].dy, s.dots[i].dy = s.dots[i].dy, s.dots[j].dy
				}
			}
		}

		gc.Color(8)
		gc.CircFill(s.dots[i].x, s.dots[i].y, s.dots[i].r)

		gc.Color(7)
		gc.Circ(s.dots[i].x, s.dots[i].y, s.dots[i].r)

		s.dots[i].x += s.dots[i].dx
		s.dots[i].y += s.dots[i].dy
	}
}

type dot struct {
	x  int
	y  int
	dx int
	dy int
	r  int
}
