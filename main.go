package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/utilyre/bb/config"
	"github.com/utilyre/bb/energy"
)

var (
	isStopped bool = true

	erg energy.Energy = energy.NewEnergy(
		config.Mass,
		config.Mass*config.Gravity*config.InitialHeight, // ΔU = mgΔh
		0,
	)
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("bb: ")

	go updater()
	pixelgl.Run(renderer)
}

func updater() {
	last := time.Now()
	for {
		time.Sleep(5 * time.Millisecond)

		// Calculates elapsed time
		dt := time.Since(last).Seconds()
		last = time.Now()

		// Won't do any calculations if is stopped
		if isStopped {
			continue
		}

		dy := -0.5*config.Gravity*math.Pow(dt, 2) + erg.Velocity()*dt // Δy = -½gΔt² + V₀Δt
		h := (erg.Potential() / (config.Mass * config.Gravity) /* Δh = ΔU / mg */) + dy

		erg.SetPotential(config.Mass * config.Gravity * h) // ΔU = mgΔh
	}
}

func renderer() {
	cfg := pixelgl.WindowConfig{
		Title:     "Bouncing Ball",
		Bounds:    pixel.R(0, 0, (config.Radius+3)*config.Scale, (config.InitialHeight+2)*config.Scale),
		Resizable: true,
		VSync:     true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	win.SetSmooth(true)

	pic, err := loadPicture(filepath.Join("assets", "basketball.png"))
	if err != nil {
		log.Fatalln(err)
	}

	basketball := pixel.NewSprite(pic, pic.Bounds())

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(color.RGBA{R: 43, G: 45, B: 66, A: 255})

		if win.JustPressed(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}

		if win.JustPressed(pixelgl.KeySpace) {
			isStopped = !isStopped
		}

		if !isStopped {
			if win.Pressed(pixelgl.KeyW) {
				erg.ExertForce(config.Force, dt)
			}

			if win.Pressed(pixelgl.KeyS) {
				erg.ExertForce(-config.Force, dt)
			}
		}

		h := erg.Potential() / (config.Mass * config.Gravity) // Δh = ΔU / mg
		basketball.Draw(
			win,
			pixel.IM.Scaled(
				pixel.ZV,
				2*config.Radius*config.Scale/102,
			).Moved(
				pixel.V(
					win.Bounds().Center().X,
					(h+config.Radius)*config.Scale,
				),
			),
		)

		win.Update()
	}
}

func loadPicture(name string) (pixel.Picture, error) {
	file, err := os.Open(name)
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
