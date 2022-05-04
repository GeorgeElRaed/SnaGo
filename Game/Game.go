package game

import (
	"fmt"
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

func handleSnakeCollision(event snake.CollisionEvent) {
	if event.CollidedWith == "snake" {
		fmt.Println("DEATHHHHHH")
	}

	if event.CollidedWith == "apple" {
		event.Snake.Grow()
		event.Snake.Apple.Reposition()
	}
}

func Create() *Game {
	g := Game{em: em.Init()}
	g.em.Add(&grid.Grid{})
	a := Apple.Apple{}
	g.em.Add(&a)
	g.em.Add(&snake.Snake{Color: colornames.Green, Apple: &a, OnCollisionListeners: []func(event snake.CollisionEvent){handleSnakeCollision}})

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
