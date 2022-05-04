package grid

import (
	"github.com/GeorgeElRaed/SnaGo/Config"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Grid struct {
	gridLines *imdraw.IMDraw
}

func (g *Grid) Init() {
	g.gridLines = imdraw.New(nil)
	g.gridLines.Color = colornames.Black
	g.gridLines.Clear()

	for x := Config.BlockWidth; x < Config.WindowWidth; x += Config.BlockWidth {
		g.gridLines.Push(pixel.V(x, 0), pixel.V(x, Config.WindowHeight))
		g.gridLines.Line(1)
	}

	for y := Config.BlockHeight; y < Config.WindowHeight; y += Config.BlockHeight {
		g.gridLines.Push(pixel.V(0, y), pixel.V(Config.WindowWidth, y))
		g.gridLines.Line(1)
	}
}

func (g *Grid) Update() {
}

func (g *Grid) Render(win *pixelgl.Window) {

	g.gridLines.Draw(win)
}
