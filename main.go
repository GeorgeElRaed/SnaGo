package main

import (
	"image"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	_ "image/png"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Something",
		Bounds: pixel.R(0, 0, 640, 480),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	pic, err := loadPicture("red.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	x := 50.0
	y := 50.0
	xD := 1.0
	yD := 1.0
	for !win.Closed() {
		win.Clear(colornames.Skyblue)
		sprite.Draw(win, pixel.IM.Scaled(pixel.Vec{X: 1, Y: 1}, 0.25).Moved(pixel.Vec{X: x, Y: y}))

		if x+sprite.Frame().W()*0.25/2 > 640 || x < sprite.Frame().W()*0.25/2 {
			xD *= -1
		}
		if y+sprite.Frame().H()*0.25/2 > 480 || y < sprite.Frame().H()*0.25/2 {
			yD *= -1
		}

		x += xD * 0.1
		y += yD * 0.1
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
