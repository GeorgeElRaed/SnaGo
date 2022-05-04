package main

import (
	"github.com/GeorgeElRaed/SnaGo/Config"
	_ "image/png"

	"github.com/GeorgeElRaed/SnaGo/game"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "SnaGo",
		Bounds: pixel.R(0, 0, Config.WindowWidth, Config.WindowHeight),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	g := game.Init()

	g.Init()
	for !win.Closed() {
		win.Clear(colornames.Skyblue)
		g.Update()
		g.Render(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
