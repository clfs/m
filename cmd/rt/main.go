package main

import (
	"image/png"
	"log"
	"os"

	"github.com/clfs/m/raytracer"
)

func main() {
	cfg := raytracer.Config{
		ImageHeight: 256,
		ImageWidth:  256,
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
