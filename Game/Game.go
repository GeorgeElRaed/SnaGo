package game

import (
	"github.com/GeorgeElRaed/SnaGo/Apple"
	snake "github.com/GeorgeElRaed/SnaGo/Snake"
	em "github.com/GeorgeElRaed/SnaGo/entitymanager"
	"github.com/GeorgeElRaed/SnaGo/grid"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Game struct {
	em *em.EntityManager
}

func Init() *Game {
	g := Game{em: em.Init()}
	g.em.Add(&grid.Grid{})
	g.em.Add(&snake.Snake{Color: colornames.Green})
	g.em.Add(&Apple.Apple{})
	return &g
}

func (g *Game) Init() {
	g.em.Init()
}

func (g *Game) Update() {
	g.em.Update()
}

func (g *Game) Render(win *pixelgl.Window) {
	g.em.Render(win)
}
