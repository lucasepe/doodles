package dotscube

import (
	"context"
	"math"

	"github.com/lucasepe/doodlekit"
)

func Scene() doodlekit.Scene {
	return &scene{}
}

type scene struct {
	dots  []dot
	time  float64
	speed float64
	w, h  int
}

func (s *scene) Init(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	s.w, s.h = gc.Width(), gc.Height()
	s.speed = 8.0
	s.dots = []dot{}

	for z := -1.0; z < 1.0; z += 0.5 {
		for y := -1.0; y < 1.0; y += 0.5 {
			for x := -1.0; x < 1.0; x += 0.5 {
				s.dots = append(s.dots, dot{
					x:  x,
					y:  y,
					z:  -z,
					ci: 9 + int(x*2+y*3)%9,
				})
			}
		}
	}
}

func (s *scene) Update(ctx context.Context, dt float64) {
	s.time = s.time + s.speed*dt
	// camera space
	for i := 0; i < len(s.dots); i++ {
		s.dots[i].rot(s.time)
	}

	// 	--sort by distance from camera
	//  --because they go out of order
	//  --slowly, doing a bubble sort
	//  --up and down 3 times is good
	//  --enough
	for pass := 1; pass < 3; pass++ {
		for i := 0; i < len(s.dots)-1; i++ {
			if s.dots[i].cz < s.dots[i+1].cz {
				s.dots[i], s.dots[i+1] = s.dots[i+1], s.dots[i]
			}
		}

		for i := len(s.dots) - 1; i < 1; i-- {
			if s.dots[i].cz < s.dots[i+1].cz {
				s.dots[i], s.dots[i+1] = s.dots[i+1], s.dots[i]
			}
		}
	}
}

func (s *scene) Draw(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	rad1 := 5 + math.Cos(s.time/4)*4

	halfW, halfH := float64(s.w/2), float64(s.h/2)
	for _, d := range s.dots {
		// --> screen space
		sx := halfW + d.cx*halfW/d.cz
		sy := halfH + d.cy*halfH/d.cz
		rad := rad1 / d.cz

		if d.cz > 0.1 { //-- clip
			gc.Color(d.ci)
			gc.CircFill(int(sx), int(sy), int(rad))
			gc.Color(8)
			gc.CircFill(int(sx+rad/3), int(sy-rad/3), int(rad/3))
		}
	}
}

type dot struct {
	x, y, z    float64
	cx, cy, cz float64
	ci         int
}

func (d *dot) rot(st float64) {
	d.cx, d.cz = rot(d.x, d.z, st/8)
	d.cy, d.cz = rot(d.y, d.cz, st/7)
	d.cz += 2 + math.Cos(st/3)
}

func rot(x, y, a float64) (float64, float64) {
	xr := math.Cos(a)*x - math.Sin(a)*y
	yr := math.Cos(a)*y + math.Sin(a)*x
	return xr, yr
}
