package Apple

import (
	"github.com/GeorgeElRaed/SnaGo/Config"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Apple struct {
	appleRenderer *imdraw.IMDraw
	position      pixel.Vec
}

func (a *Apple) Init() {
	a.appleRenderer = imdraw.New(nil)
	a.appleRenderer.Color = colornames.Red
	a.appleRenderer.Clear()

	a.position = pixel.V(4, 4)

	a.appleRenderer.Push(pixel.V(a.position.X*Config.BlockWidth, a.position.Y*Config.BlockHeight), pixel.V(a.position.X*Config.BlockWidth+Config.BlockWidth, a.position.Y*Config.BlockHeight+Config.BlockHeight))
	a.appleRenderer.Rectangle(0)
	a.appleRenderer.Push(pixel.V(a.position.X*Config.BlockWidth, a.position.Y*Config.BlockHeight), pixel.V(a.position.X*Config.BlockWidth+Config.BlockWidth, a.position.Y*Config.BlockHeight+Config.BlockHeight))
	a.appleRenderer.Rectangle(1)

}

func (a *Apple) Update() {
}

func (a *Apple) Render(win *pixelgl.Window) {
	a.appleRenderer.Draw(win)
}
