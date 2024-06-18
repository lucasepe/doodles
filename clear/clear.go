package clear

import (
	"context"

	"github.com/lucasepe/doodlekit"
)

func Scene(color int) doodlekit.Scene {
	return &scene{
		color: color,
	}
}

type scene struct {
	color int
}

func (s *scene) Init(ctx context.Context) {
	doodlekit.Canvas(ctx).Cls(s.color)
}

func (s *scene) Update(ctx context.Context, dt float64) {}

func (s *scene) Draw(ctx context.Context) {
	doodlekit.Canvas(ctx).Cls(s.color)
}
