package starfield

import (
	"context"
	"math"

	"github.com/lucasepe/doodlekit"
)

func Scene(roto bool) doodlekit.Scene {
	return &scene{
		roto:  roto,
		speed: 1.5,
		total: 250,
	}
}

type scene struct {
	w, h   int
	speed  float64
	stars  []star
	total  int
	depth  float64
	roto   bool
	angleZ float64
}

func (s *scene) Init(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)
	rng := doodlekit.Rng(ctx)

	s.w, s.h = gc.Width(), gc.Height()
	s.depth = 0.7 * float64(s.w)

	s.stars = make([]star, s.total)
	for i := 0; i < len(s.stars); i++ {
		s.stars[i].x = rng.Rnd(-0.5*float64(s.w), 0.5*float64(s.w))
		s.stars[i].y = rng.Rnd(-0.5*float64(s.h), 0.5*float64(s.h))
		s.stars[i].z = rng.Rnd(1, s.depth)
	}
}

func (s *scene) Update(ctx context.Context, dt float64) {
	rng := doodlekit.Rng(ctx)

	cos, sin := math.Cos, math.Sin

	for i := 0; i < s.total; i++ {
		// Move star towards the viewer
		s.stars[i].z -= s.speed
		if s.stars[i].z <= 0 {
			s.stars[i].x = rng.Rnd(-0.5*float64(s.w), 0.5*float64(s.w))
			s.stars[i].y = rng.Rnd(-0.5*float64(s.h), 0.5*float64(s.h))
			s.stars[i].z = s.depth
		}

		if s.roto {
			// Rotate around Z axis
			x_rot := s.stars[i].x*cos(s.angleZ) - s.stars[i].y*sin(s.angleZ)
			y_rot := s.stars[i].x*sin(s.angleZ) + s.stars[i].y*cos(s.angleZ)
			s.stars[i].x = x_rot
			s.stars[i].y = y_rot
		}
	}

	if s.roto {
		s.angleZ += 0.0001
	}
}

func (s *scene) Draw(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)
	gc.Translate(0.5*float64(s.w), 0.5*float64(s.h))

	gc.Color(6) //int(s.stars[i].z))

	for i := 0; i < s.total; i++ {
		// Convert 3D position to 2D screen position
		screen_x := (s.stars[i].x / s.stars[i].z) * s.depth
		screen_y := (s.stars[i].y / s.stars[i].z) * s.depth

		gc.Pix(int(screen_x), int(screen_y))
	}

	gc.Identity()
}

type star struct {
	x float64
	y float64
	z float64
}
