package connections

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
	w, h  int
	total int
	dots  []dot
}

func (s *scene) Init(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)
	rng := doodlekit.Rng(ctx)

	s.w, s.h = gc.Width(), gc.Height()

	if s.total <= 0 || s.total > 30 {
		s.total = 16
	}

	s.dots = make([]dot, s.total)
	for i := 0; i < s.total; i++ {
		r := rng.RndI(2, 3)
		s.dots[i] = dot{
			x:  rng.RndI(r, s.w-r),
			y:  rng.RndI(r, s.h-r),
			dx: rng.RndI(1, 2),
			dy: rng.RndI(1, 2),
			r:  r,
		}
	}
}

func (s *scene) Update(ctx context.Context, dt float64) {
	for i := 0; i < s.total; i++ {
		r := s.dots[i].r
		xc := s.dots[i].x
		yc := s.dots[i].y

		if xc-r <= r || xc+r >= s.w-r {
			s.dots[i].dx = -s.dots[i].dx
		}

		if yc-r <= r || yc+r >= s.h-r {
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
				gc.Color(12)
				gc.Line(s.dots[j].x, s.dots[j].y, s.dots[i].x, s.dots[i].y)
			}

			if d < dr {
				s.dots[j].dx, s.dots[i].dx = s.dots[i].dx, s.dots[j].dx
				s.dots[j].dy, s.dots[i].dy = s.dots[i].dy, s.dots[j].dy
			}
		}

		gc.Color(9)
		gc.CircFill(s.dots[i].x, s.dots[i].y, s.dots[i].r)

		gc.Color(8)
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
