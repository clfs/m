package main

import (
	"image/png"
	"log"
	"os"

	"github.com/clfs/m/math/f64"
	"github.com/clfs/m/raytracer"
)

func main() {
	cfg := raytracer.Config{
		AspectRatio:    16.0 / 9.0,
		ImageWidth:     400,
		FocalLength:    1.0,
		ViewportHeight: 2.0,
		CameraCenter:   f64.Vec3{0, 0, 0},
	}

	img := raytracer.Render(cfg)

	f, err := os.Create("demo.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}
