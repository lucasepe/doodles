package doodles

import (
	"testing"

	"github.com/lucasepe/doodlekit"
	"github.com/lucasepe/doodles/clear"
	"github.com/lucasepe/doodles/clock"
	"github.com/lucasepe/doodles/connections"
	"github.com/lucasepe/doodles/glitters"
	"github.com/lucasepe/doodles/qix"
	"github.com/lucasepe/doodles/sinebobs"
	"github.com/lucasepe/doodles/starfield"
	"github.com/lucasepe/doodles/stars"
	"github.com/lucasepe/doodles/tunnel"
	"github.com/lucasepe/doodles/xmastree"
)

func TestClockScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(0),
		starfield.Scene(true),
		clock.Scene(),
	}

	doodlekit.NewLoop("clock",
		//doodlekit.X2(),
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(20),
	).Run(scenes)
}

func TestStarfield3dRotoScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(0),
		starfield.Scene(true),
	}

	doodlekit.NewLoop("starfield-roto",
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestStarfield3dScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(0),
		starfield.Scene(false),
	}

	doodlekit.NewLoop("starfield",
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestSinebobsScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(0),
		stars.Scene(),
		sinebobs.Scene(),
	}

	doodlekit.NewLoop("sinebobs",
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestXmasTreeWithGlittersScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(0),
		glitters.Scene(),
		xmastree.Scene(),
	}

	doodlekit.NewLoop("glitters-xmastree",
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestGlittersScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(0),
		glitters.Scene(),
	}

	doodlekit.NewLoop("glitters",
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestXmasTreeScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(0),
		xmastree.Scene(),
	}

	doodlekit.NewLoop("xmastree",
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestConnectionsScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(0),
		connections.Scene(30),
	}

	doodlekit.NewLoop("connections",
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestQixScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(0),
		qix.Scene(),
	}

	doodlekit.NewLoop("qix",
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

func TestStarsScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(0),
		stars.Scene(),
	}

	doodlekit.NewLoop("stars",
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}

// go.testTimeout 90s
func TestTunnelScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		tunnel.Scene(),
	}

	doodlekit.NewLoop("tunnel",
		doodlekit.OutDir("_doodles"),
		doodlekit.StopAfter(15),
	).Run(scenes)
}
