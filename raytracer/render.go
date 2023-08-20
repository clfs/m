package raytracer

import (
	"image"
	"image/color"
	"log"
)

var Logger = log.Default()

type Config struct {
	ImageHeight int
	ImageWidth  int
}

func Render(c Config) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, c.ImageWidth, c.ImageHeight))
	for j := 0; j < c.ImageHeight; j++ {
		log.Printf("Scanlines remaining: %d", c.ImageHeight-j)
		for i := 0; i < c.ImageWidth; i++ {
			r := float64(i) / float64(c.ImageWidth-1)
			g := float64(j) / float64(c.ImageHeight-1)
			b := 0.0
			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)
			img.SetRGBA(i, j, color.RGBA{uint8(ir), uint8(ig), uint8(ib), 255})
		}
	}
	return img
}
