package snake

import (
	"github.com/GeorgeElRaed/SnaGo/Apple"
	"github.com/GeorgeElRaed/SnaGo/Config"
	"github.com/GeorgeElRaed/SnaGo/InputMap"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

type CollisionEvent struct {
	CollidedWith string
	Snake        *Snake
}

type Snake struct {
	Color                color.RGBA
	snakeRenderer        *imdraw.IMDraw
	snakeParts           []pixel.Vec
	moveDirection        pixel.Vec
	OnCollisionListeners []func(CollisionEvent)
	Apple                *Apple.Apple
}

func (s *Snake) Init() {
	s.snakeParts = make([]pixel.Vec, 0)
	s.snakeParts = append(s.snakeParts, pixel.V(3, 1), pixel.V(2, 1), pixel.V(1, 1))

	s.moveDirection = pixel.V(1, 0)

	s.snakeRenderer = imdraw.New(nil)
	s.snakeRenderer.Color = s.Color

}

func (s *Snake) Update() {

	if InputMap.Inputs.Contains(pixelgl.KeyLeft) && s.moveDirection.Y != 0 {
		s.moveDirection = pixel.V(-1, 0)
	} else if InputMap.Inputs.Contains(pixelgl.KeyRight) && s.moveDirection.Y != 0 {
		s.moveDirection = pixel.V(1, 0)
	} else if InputMap.Inputs.Contains(pixelgl.KeyUp) && s.moveDirection.X != 0 {
		s.moveDirection = pixel.V(0, 1)
	} else if InputMap.Inputs.Contains(pixelgl.KeyDown) && s.moveDirection.X != 0 {
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

	for i, part := range s.snakeParts {
		if i == 0 {
			continue
		}
		if s.snakeParts[0].Eq(part) {
			s.collidedWith("snake")
		}
	}

	if s.snakeParts[0].Eq(s.Apple.Position) {
		s.collidedWith("apple")
	}

}

func (s *Snake) collidedWith(with string) {
	for _, handler := range s.OnCollisionListeners {
		handler(CollisionEvent{CollidedWith: with, Snake: s})
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

func (s *Snake) Grow() {
	s.snakeParts = append(s.snakeParts, s.snakeParts[len(s.snakeParts)-1])
}
