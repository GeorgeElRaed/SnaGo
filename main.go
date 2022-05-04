package main

import (
	"github.com/GeorgeElRaed/SnaGo/Config"
	"github.com/GeorgeElRaed/SnaGo/InputMap"
	_ "image/png"
	"time"

	"github.com/GeorgeElRaed/SnaGo/game"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func getCurrentTime() float64 {
	return float64(time.Now().UnixNano()) * 1e-9
}

func processInputs(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyUp) {
		InputMap.Inputs.Add(pixelgl.KeyUp)
	}
	if win.Pressed(pixelgl.KeyDown) {
		InputMap.Inputs.Add(pixelgl.KeyDown)
	}
	if win.Pressed(pixelgl.KeyLeft) {
		InputMap.Inputs.Add(pixelgl.KeyLeft)
	}
	if win.Pressed(pixelgl.KeyRight) {
		InputMap.Inputs.Add(pixelgl.KeyRight)
	}

}

func clearInputs() {
	InputMap.Inputs.Clear()
}

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

	const idleThreshold = 1
	var current, elapsed, previous, lag float64

	previous = getCurrentTime()
	lag = 0.0
	for !win.Closed() {

		current = getCurrentTime()
		elapsed = current - previous
		previous = current

		if elapsed > idleThreshold {
			win.Update()
			continue
		}
		lag += elapsed

		processInputs(win)
		updated := false
		for lag >= Config.SecondsPerUpdate {
			g.Update()
			lag -= Config.SecondsPerUpdate
			updated = true
		}
		if updated {
			clearInputs()
		}

		win.Clear(colornames.Skyblue)
		g.Render(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
