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
	"github.com/utilyre/bb/config"
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

func worker(ch chan<- energy.Energy) {
	energy := energy.NewEnergy(config.Mass*config.Gravity*config.InitialHeight, 0)

	last := time.Now()
	for {
		dt := time.Since(last).Seconds()
		last = time.Now()

		coefficient := 1.0
		if energy.IsFalling() {
			coefficient = -1.0
		}

		dx := config.Gravity*math.Pow(dt, 2)/2 + energy.Speed()*dt
		h := energy.Potential()/(config.Mass*config.Gravity) + coefficient*dx
		energy.SetPotential(config.Mass * config.Gravity * h)

		ch <- energy
	}
}

func run(ch <-chan energy.Energy) {
	cfg := pixelgl.WindowConfig{
		Title:  "Bouncing Ball",
		Bounds: pixel.R(0, 0, 2*config.Scale, (config.InitialHeight+1)*config.Scale),
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
			h := (energy.Potential() / (config.Mass * config.Gravity)) * config.Scale

			basketball.Draw(
				win,
				pixel.IM.Scaled(
					pixel.ZV,
					2*config.Radius*config.Scale/128,
				).Moved(
					pixel.V(
						win.Bounds().Center().X,
						h+config.Radius*config.Scale,
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
