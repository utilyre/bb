package main

import (
	"image/color"
	"log"
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/utilyre/bb/me"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("bb: ")

	ch := make(chan me.MechanicalEnergy)
	defer close(ch)

	go worker(ch)
	pixelgl.Run(func() { run(ch) })
}

const (
	mass          = 0.5 // Kg
	gravity       = 9.8 // m/s^2
	initialHeight = 5.0 // m

	scale = 100.0 // px/m
)

func worker(ch chan<- me.MechanicalEnergy) {
	energy := me.New(mass*gravity*initialHeight, 0)

	last := time.Now()
	for {
		dt := time.Since(last).Seconds()
		last = time.Now()

		dx := gravity*math.Pow(dt, 2)/2 + energy.Speed()*dt // x = 1/2at^2 + v0t
		h := energy.Potential()/(mass*gravity) - dx         // h = U/(mg)

		if h > 0 {
			energy.SetPotential(energy, mass*gravity*h) // U = mgh
		} else {
			energy.SetSpeed(energy, -energy.Speed()) // v = -v0
		}

		ch <- energy
	}
}

func run(ch <-chan me.MechanicalEnergy) {
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

	ball := imdraw.New(nil)
	ball.Color = color.RGBA{R: 38, G: 70, B: 83, A: 255}

	for !win.Closed() {
		win.Clear(color.RGBA{R: 42, G: 157, B: 143, A: 255})
		ball.Clear()

		if energy, ok := <-ch; ok {
			h := (energy.Potential() / (mass * gravity)) * scale
			ball.Push(pixel.V(win.Bounds().Center().X, h))
		}
		ball.Circle(10, 0)

		ball.Draw(win)
		win.Update()
	}
}
