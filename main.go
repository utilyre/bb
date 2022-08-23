package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/utilyre/bb/energy"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("bb: ")

	ch := make(chan energy.Energy)
	defer close(ch)

	go worker(ch)
	pixelgl.Run(func() { run(ch) })
}

const (
	radius  = 0.25 // m
	mass    = 0.5  // kg
	gravity = 9.8  // m/s^2

	initialHeight = 2.0 // m
	energyLoss    = 2.0 // j

	scale = 200.0 // px/m
)

func worker(ch chan<- energy.Energy) {
	energy := energy.NewEnergy(mass*gravity*initialHeight, 0)

	last := time.Now()
	for {
		dt := time.Since(last).Seconds()
		last = time.Now()

		multiplier := 1.0
		if energy.IsFalling() {
			multiplier = -1.0
		}

		dx := gravity*math.Pow(dt, 2)/2 + energy.Speed()*dt    // x = 1/2at^2 + v0t
		h := energy.Potential()/(mass*gravity) + multiplier*dx // h = U/(mg)
		energy.SetPotential(mass * gravity * h)                // U = mgh

		ch <- energy
	}
}

func run(ch <-chan energy.Energy) {
	cfg := pixelgl.WindowConfig{
		Title:  "Bouncing Ball",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	win.SetSmooth(true)

	pic, err := loadPicture("assets/basketball.png")
	if err != nil {
		log.Fatalln(err)
	}

	basketball := pixel.NewSprite(pic, pic.Bounds())

	for !win.Closed() {
		win.Clear(color.RGBA{R: 43, G: 45, B: 66, A: 255})

		if energy, ok := <-ch; ok {
			h := (energy.Potential() / (mass * gravity)) * scale

			basketball.Draw(
				win,
				pixel.IM.Scaled(
					pixel.ZV,
					2*radius*scale/128,
				).Moved(
					pixel.V(
						win.Bounds().Center().X,
						h+radius*scale,
					),
				),
			)
		}

		win.Update()
	}
}

func loadPicture(filename string) (pixel.Picture, error) {
	file, err := os.Open(filename)
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
