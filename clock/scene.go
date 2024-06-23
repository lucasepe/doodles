package clock

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/lucasepe/doodlekit"
)

func Scene() doodlekit.Scene {
	return &scene{}
}

type scene struct {
	centerX int
	centerY int
	radius  int
	xH, yH  int
	xM, yM  int
	xS, yS  int
}

func (s *scene) Init(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	s.centerX, s.centerY = gc.Width()/2, gc.Height()/2
	s.radius = int(math.Round(float64(gc.Width()/2) * 0.8))
}

func (s *scene) Update(ctx context.Context, dt float64) {
	t := time.Now()

	// Calculate angles for the hour, minute, and second hands
	angleH := 2*math.Pi/12*float64(t.Hour()%12) + 2*math.Pi/12/60*float64(t.Minute())
	angleM := 2 * math.Pi / 60 * float64(t.Minute())
	angleS := 2 * math.Pi / 60 * float64(t.Second())

	length := float64(s.radius) * 0.5
	s.xH = s.centerX + int(length*math.Cos(angleH-math.Pi/2))
	s.yH = s.centerY + int(length*math.Sin(angleH-math.Pi/2))

	length = float64(s.radius) * 0.65
	s.xM = s.centerX + int(length*math.Cos(angleM-math.Pi/2))
	s.yM = s.centerY + int(length*math.Sin(angleM-math.Pi/2))

	length = float64(s.radius) * 0.75
	s.xS = s.centerX + int(length*math.Cos(angleS-math.Pi/2))
	s.yS = s.centerY + int(length*math.Sin(angleS-math.Pi/2))
}

func (s *scene) Draw(ctx context.Context) {
	gc := doodlekit.Canvas(ctx)

	//gc.Color(4)
	//gc.CircFill(s.centerX, s.centerY, s.radius)

	//gc.Color(7)
	//gc.CircFill(s.centerX, s.centerY, s.radius-3)

	gc.Color(2)

	r := 1.0 + float64(s.radius)*0.8
	for hour := 1; hour <= 12; hour++ {
		angle := 2 * math.Pi / 12 * float64(hour-3) // -3 per mettere l'ora 12 in alto
		x := s.centerX + int(r*math.Cos(angle))
		y := s.centerY + int(r*math.Sin(angle))

		lbl := fmt.Sprintf("%d", hour)
		tw, th := gc.MeasureString(lbl)
		x -= tw / 2
		y += th / 2
		gc.Print(lbl, x, y)
	}

	gc.Color(5)
	gc.Line(s.centerX, s.centerY, s.xH, s.yH)

	gc.Color(13)
	gc.Line(s.centerX, s.centerY, s.xM, s.yM)

	gc.Color(7)
	gc.Line(s.centerX, s.centerY, s.xS, s.yS)

}
