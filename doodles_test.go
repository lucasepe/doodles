package doodles

import (
	"testing"

	"github.com/lucasepe/doodlekit"
	"github.com/lucasepe/doodles/clear"
	"github.com/lucasepe/doodles/clock"
	"github.com/lucasepe/doodles/connections"
	"github.com/lucasepe/doodles/glitters"
	"github.com/lucasepe/doodles/lines"
	"github.com/lucasepe/doodles/sinebobs"
	"github.com/lucasepe/doodles/starfield"
	"github.com/lucasepe/doodles/stars"
	"github.com/lucasepe/doodles/tunnel"
	"github.com/lucasepe/doodles/xmastree"
)

const (
	scaleFactor = 2
)

func TestClockWithStarfieldScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(1),
		starfield.Scene(true),
		clock.Scene(),
	}

	doodlekit.NewLoop("clock-starfield",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(20),
	).Run(scenes)
}

func TestClockScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(1),
		clock.Scene(),
	}

	doodlekit.NewLoop("clock",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(20),
	).Run(scenes)
}

func TestStarfield3dRotoScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(1),
		starfield.Scene(true),
	}

	doodlekit.NewLoop("starfield-roto",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestStarfield3dScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(1),
		starfield.Scene(false),
	}

	doodlekit.NewLoop("starfield",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestSinebobsScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(1),
		sinebobs.Scene(),
	}

	doodlekit.NewLoop("sinebobs",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestXmasTreeWithGlittersScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(1),
		glitters.Scene(50),
		xmastree.Scene(),
	}

	doodlekit.NewLoop("glitters-xmastree",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestGlittersScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(1),
		glitters.Scene(80),
	}

	doodlekit.NewLoop("glitters",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestXmasTreeScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(1),
		xmastree.Scene(),
	}

	doodlekit.NewLoop("xmastree",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestConnectionsScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(1),
		connections.Scene(8),
	}

	doodlekit.NewLoop("connections",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(10),
	).Run(scenes)
}

func TestConnectionsWithStarsScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(1),
		stars.Scene(120),
		connections.Scene(8),
	}

	doodlekit.NewLoop("connections-stars",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(10),
	).Run(scenes)
}

func TestLinesScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(1),
		lines.Scene(),
	}

	doodlekit.NewLoop("lines",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(10),
	).Run(scenes)
}

func TestLinesWithStarfieldScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(1),
		starfield.Scene(true),
		lines.Scene(),
	}

	doodlekit.NewLoop("lines-starfield",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(10),
	).Run(scenes)
}

func TestStarsScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(1),
		stars.Scene(128),
	}

	doodlekit.NewLoop("stars",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestTunnelScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		tunnel.Scene(),
	}

	doodlekit.NewLoop("tunnel",
		doodlekit.Resize(scaleFactor),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}
