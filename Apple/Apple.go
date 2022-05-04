package Apple

import (
	"github.com/GeorgeElRaed/SnaGo/Config"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math/rand"
)

type Apple struct {
	appleRenderer *imdraw.IMDraw
	Position      pixel.Vec
}

func (a *Apple) Init() {
	a.appleRenderer = imdraw.New(nil)
	a.appleRenderer.Color = colornames.Red

	a.Position = pixel.V(4, 4)
}

func (a *Apple) Update() {
}

func (a *Apple) Render(win *pixelgl.Window) {
	a.appleRenderer.Clear()

	a.appleRenderer.Push(pixel.V(a.Position.X*Config.BlockWidth, a.Position.Y*Config.BlockHeight), pixel.V(a.Position.X*Config.BlockWidth+Config.BlockWidth, a.Position.Y*Config.BlockHeight+Config.BlockHeight))
	a.appleRenderer.Rectangle(0)
	a.appleRenderer.Push(pixel.V(a.Position.X*Config.BlockWidth, a.Position.Y*Config.BlockHeight), pixel.V(a.Position.X*Config.BlockWidth+Config.BlockWidth, a.Position.Y*Config.BlockHeight+Config.BlockHeight))
	a.appleRenderer.Rectangle(1)
	a.appleRenderer.Draw(win)
}

func (a *Apple) Reposition() {
	randomX := rand.Intn(int(Config.GridHorizontalCount))
	randomY := rand.Intn(int(Config.GridVerticalCount))
	a.Position = pixel.V(float64(randomX), float64(randomY))
}
