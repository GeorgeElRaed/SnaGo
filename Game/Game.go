package game

import (
	"github.com/GeorgeElRaed/SnaGo/Apple"
	grid "github.com/GeorgeElRaed/SnaGo/Grid"
	snake "github.com/GeorgeElRaed/SnaGo/Snake"
	em "github.com/GeorgeElRaed/SnaGo/entitymanager"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Game struct {
	em         *em.EntityManager
	OnGameOver func(*Game)
}

func handleSnakeCollision(event snake.CollisionEvent, game *Game) {
	if event.CollidedWith == "snake" {
		game.OnGameOver(game)
	}

	if event.CollidedWith == "apple" {
		event.Snake.Grow()
		event.Snake.Apple.Reposition()
	}
}

func Create(OnGameOver func(*Game)) *Game {
	g := Game{OnGameOver: OnGameOver}
	return &g
}

func (g *Game) Init() {
	g.em = em.Create()
	g.em.Add(&grid.Grid{})
	a := Apple.Apple{}
	g.em.Add(&a)
	g.em.Add(&snake.Snake{Color: colornames.Green, Apple: &a, OnCollisionListeners: []func(event snake.CollisionEvent){
		func(event snake.CollisionEvent) {
			handleSnakeCollision(event, g)
		},
	}})

	g.em.Init()
}

func (g *Game) Update() {
	g.em.Update()
}

func (g *Game) Render(win *pixelgl.Window) {
	g.em.Render(win)
}
