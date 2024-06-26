package sinebobs

import (
	"context"
	"math"

	"github.com/lucasepe/doodlekit"
)

func Scene() doodlekit.Scene {
	return &scene{}
}

type scene struct {
	w, h               int
	time               float64
	cols, rows         int
	spacingX, spacingY int
	offsetX, offsetY   int
}

func (s *scene) Init(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	s.w, s.h = gc.Width(), gc.Height()
	// Number of spheres in each direction
	s.cols, s.rows = 8, 6
	// Small margin to avoid cutting spheres at the edges
	colsMargin, rowsMargin := 4, 10
	// Spacing between spheres
	s.spacingX = (s.w - 2*colsMargin) / (s.cols - 1)
	s.spacingY = (s.h - 2*rowsMargin) / (s.rows - 1)

	s.offsetX = (s.w - (s.cols-1)*s.spacingX) / 2
	s.offsetY = (s.h - (s.rows-1)*s.spacingY) / 2
}

func (s *scene) Update(ctx context.Context, dt float64) {
	s.time += 4 * dt
}

func (s *scene) Draw(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	for l := 0.0; l < float64(s.cols); l++ {
		x := int(l*float64(s.spacingX)) + s.offsetX

		for n := 0.0; n < float64(s.rows); n++ {
			y := int(n*float64(s.spacingY)) + s.offsetY
			p := math.Sin(s.time + l/2 - n/1.5*math.Sin(s.time+n/8))

			gc.Color(3 + int(s.time))
			gc.CircFill(x, y, int(p+2))

			gc.Color(9 + int(s.time))
			gc.CircFill(x, y, int(p+1))
		}
	}
}
