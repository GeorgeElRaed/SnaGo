package snake

import (
	"github.com/GeorgeElRaed/SnaGo/Config"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

type SnakePart struct {
	x, y float64
}

type Snake struct {
	Color         color.RGBA
	snakeRenderer *imdraw.IMDraw
	snakeParts    []pixel.Vec
}

func (s *Snake) Init() {
	s.snakeParts = make([]pixel.Vec, 0)
	s.snakeParts = append(s.snakeParts, pixel.V(1, 1), pixel.V(2, 1))

	s.snakeRenderer = imdraw.New(nil)
	s.snakeRenderer.Color = s.Color
	s.snakeRenderer.Clear()

	for _, part := range s.snakeParts {
		s.snakeRenderer.Push(pixel.V(part.X*Config.BlockWidth, part.Y*Config.BlockHeight), pixel.V(part.X*Config.BlockWidth+Config.BlockWidth, part.Y*Config.BlockHeight+Config.BlockHeight))
		s.snakeRenderer.Rectangle(0)
		s.snakeRenderer.Push(pixel.V(part.X*Config.BlockWidth, part.Y*Config.BlockHeight), pixel.V(part.X*Config.BlockWidth+Config.BlockWidth, part.Y*Config.BlockHeight+Config.BlockHeight))
		s.snakeRenderer.Rectangle(1)
	}

}

func (s *Snake) Update() {
}

func (s *Snake) Render(win *pixelgl.Window) {
	s.snakeRenderer.Draw(win)
}
