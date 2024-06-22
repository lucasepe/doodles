package doodles

import (
	"testing"

	"github.com/lucasepe/doodlekit"
	"github.com/lucasepe/doodles/clear"
	"github.com/lucasepe/doodles/connections"
	"github.com/lucasepe/doodles/glitters"
	"github.com/lucasepe/doodles/qix"
	"github.com/lucasepe/doodles/sinebobs"
	"github.com/lucasepe/doodles/starfield"
	"github.com/lucasepe/doodles/stars"
	"github.com/lucasepe/doodles/tunnel"
	"github.com/lucasepe/doodles/xmastree"
)

func TestStarfield3DScene(t *testing.T) {
	scenes := []doodlekit.Scene{
		clear.Scene(0),
		starfield.Scene(),
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
		connections.Scene(),
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
