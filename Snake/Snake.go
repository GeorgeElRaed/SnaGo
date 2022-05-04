package snake

import (
	"github.com/GeorgeElRaed/SnaGo/Config"
	"github.com/GeorgeElRaed/SnaGo/InputMap"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

type Snake struct {
	Color         color.RGBA
	snakeRenderer *imdraw.IMDraw
	snakeParts    []pixel.Vec
	moveDirection pixel.Vec
}

func (s *Snake) Init() {
	s.snakeParts = make([]pixel.Vec, 0)
	s.snakeParts = append(s.snakeParts, pixel.V(3, 1), pixel.V(2, 1), pixel.V(1, 1))

	s.moveDirection = pixel.V(1, 0)

	s.snakeRenderer = imdraw.New(nil)
	s.snakeRenderer.Color = s.Color

}
func (s *Snake) Update() {

	if InputMap.Inputs.Contains(pixelgl.KeyLeft) {
		s.moveDirection = pixel.V(-1, 0)
	}
	if InputMap.Inputs.Contains(pixelgl.KeyRight) {
		s.moveDirection = pixel.V(1, 0)
	}
	if InputMap.Inputs.Contains(pixelgl.KeyUp) {
		s.moveDirection = pixel.V(0, 1)
	}
	if InputMap.Inputs.Contains(pixelgl.KeyDown) {
		s.moveDirection = pixel.V(0, -1)
	}
	for i := range s.snakeParts {
		if i < len(s.snakeParts)-1 {
			s.snakeParts[len(s.snakeParts)-i-1] = s.snakeParts[len(s.snakeParts)-i-2]
		}
	}
	s.snakeParts[0] = s.snakeParts[0].Add(s.moveDirection)
	if s.snakeParts[0].X < 0 {
		s.snakeParts[0] = pixel.V(Config.GridHorizontalCount-1, s.snakeParts[0].Y)
	}
	if s.snakeParts[0].X >= Config.GridHorizontalCount {
		s.snakeParts[0] = pixel.V(0, s.snakeParts[0].Y)
	}
	if s.snakeParts[0].Y < 0 {
		s.snakeParts[0] = pixel.V(s.snakeParts[0].X, Config.GridVerticalCount-1)
	}
	if s.snakeParts[0].Y >= Config.GridVerticalCount {
		s.snakeParts[0] = pixel.V(s.snakeParts[0].X, 0)
	}
}

func (s *Snake) Render(win *pixelgl.Window) {
	s.snakeRenderer.Clear()
	for _, part := range s.snakeParts {
		s.snakeRenderer.Push(pixel.V(part.X*Config.BlockWidth, part.Y*Config.BlockHeight), pixel.V(part.X*Config.BlockWidth+Config.BlockWidth, part.Y*Config.BlockHeight+Config.BlockHeight))
		s.snakeRenderer.Rectangle(0)
		s.snakeRenderer.Push(pixel.V(part.X*Config.BlockWidth, part.Y*Config.BlockHeight), pixel.V(part.X*Config.BlockWidth+Config.BlockWidth, part.Y*Config.BlockHeight+Config.BlockHeight))
		s.snakeRenderer.Rectangle(1)
	}
	s.snakeRenderer.Draw(win)
}
