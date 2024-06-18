package qix

import (
	"context"

	"github.com/lucasepe/doodlekit"
)

func Scene() doodlekit.Scene {
	return &scene{
		xMin: 0,
		yMin: 0,
		xMax: 160,
		yMax: 160,
	}
}

type scene struct {
	deck         []line
	color        int
	dx1, dy1     int
	dx2, dy2     int
	x1dir, y1dir int
	x2dir, y2dir int
	xMin, yMin   int
	xMax, yMax   int
}

func (s *scene) Init(ctx context.Context) {
	rng := doodlekit.Rng(ctx)

	s.deck = []line{}
	s.color = 1 + rng.RndI(14)

	x1, y1 := s.xMin+rng.RndI(s.xMax), s.yMin+rng.RndI(s.yMax)
	x2, y2 := s.xMin+rng.RndI(100), s.yMin+rng.RndI(100)
	for i := 0; i < 12; i++ {
		s.deck = append(s.deck, line{
			x1: x1, y1: y1,
			x2: x2, y2: y2,
			color: s.color,
		})
	}

	s.dx1 = rng.RndI(8)
	s.dx2 = rng.RndI(8)
	s.dy1 = rng.RndI(8)
	s.dy2 = rng.RndI(8)
	s.x1dir = chooseDir(ctx)
	s.x2dir = chooseDir(ctx)
	s.y1dir = chooseDir(ctx)
	s.y2dir = chooseDir(ctx)
}

func (s *scene) Update(ctx context.Context, dt float64) {
	rng := doodlekit.Rng(ctx)

	last := len(s.deck) - 1
	x1 := s.deck[last].x1 + s.dx1*s.x1dir
	x2 := s.deck[last].x2 + s.dx2*s.x2dir
	y1 := s.deck[last].y1 + s.dy1*s.y1dir
	y2 := s.deck[last].y2 + s.dy2*s.y2dir
	if x1 < s.xMin || x1 > s.xMax {
		if x1 < s.xMin {
			x1 = s.xMin
		}
		if x1 > s.xMax {
			x1 = s.xMax
		}
		s.x1dir = -s.x1dir
		s.dx1 = rng.RndI(8)
		s.color = 1 + rng.RndI(14)
	}

	if x2 < s.xMin || x2 > s.xMax {
		if x2 < s.xMin {
			x2 = s.xMin
		}
		if x2 > s.xMax {
			x2 = s.xMax
		}
		s.x2dir = -s.x2dir
		s.dx2 = rng.RndI(8)
		s.color = 1 + rng.RndI(14)
	}

	if y1 < s.yMin || y1 > s.yMax {
		if y1 < s.yMin {
			y1 = s.yMin
		}
		if y1 > s.yMax {
			y1 = s.yMax
		}
		s.y1dir = -s.y1dir
		s.dy1 = rng.RndI(8)
		s.color = 1 + rng.RndI(14)
	}

	if y2 < s.yMin || y2 > s.yMax {
		if y2 < s.yMin {
			y2 = s.yMin
		}
		if y2 > s.yMax {
			y2 = s.yMax
		}
		s.y2dir = -s.y2dir
		s.dy2 = rng.RndI(8)
		s.color = 1 + rng.RndI(14)
	}

	s.deck = append(s.deck[:0], s.deck[1:]...)
	s.deck = append(s.deck, line{
		x1: x1, y1: y1,
		x2: x2, y2: y2,
		color: s.color,
	})
}

func (s *scene) Draw(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)
	//gc.Cls(0)

	gc.Color(s.color)

	for _, el := range s.deck {
		gc.Line(el.x1, el.y1, el.x2, el.y2)
		gc.Line(s.xMax-el.y1, el.x1, s.yMax-el.y2, el.x2)
		gc.Line(el.y1, s.yMax-el.x1, el.y2, s.yMax-el.x2)
		gc.Line(s.xMax-el.x1, s.yMax-el.y1, s.xMax-el.x2, s.yMax-el.y2)
	}

}

type line struct {
	x1, y1 int
	x2, y2 int
	color  int
}

func chooseDir(ctx context.Context) int {
	rng := doodlekit.Rng(ctx)
	if rng.Rnd(1) <= 0.5 {
		return 1
	}
	return -1
}
