package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	var height, width = 256, 256

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			r := float64(i) / float64(width-1)
			g := float64(j) / float64(height-1)
			b := 0.0
			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)
			img.SetRGBA(i, j, color.RGBA{uint8(ir), uint8(ig), uint8(ib), 255})
		}
	}

	f, err := os.Create("demo.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}
