package raytracer

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/clfs/m/math/f64"
)

var Logger = log.Default()

type Config struct {
	AspectRatio    float64
	ImageWidth     int
	FocalLength    float64
	ViewportHeight float64
	CameraCenter   f64.Vec3
}

// rgbaFrom converts a vector in the unit cube to an RGBA color.
// Alpha is fixed at 255.
func rgbaFrom(v f64.Vec3) color.RGBA {
	return color.RGBA{
		R: uint8(255.999 * v[0]),
		G: uint8(255.999 * v[1]),
		B: uint8(255.999 * v[2]),
		A: 255,
	}
}

func rayColor(world Hittable, ray Ray) color.RGBA {
	var rec HitRecord
	if world.Hit(&rec, ray, 0, math.Inf(1)) {
		// = 0.5 * (<X, Y, Z> + <1, 1, 1>)
		return rgbaFrom(rec.Normal.Add(f64.Vec3{1, 1, 1}).SMul(0.5))
	}

	unitDirection := ray.Direction.Unit()
	a := 0.5 * (unitDirection[1] + 1)

	// = (1 - a) * <1, 1, 1> + a * <0.5, 0.7, 1>
	v := f64.Vec3{1, 1, 1}.SMul(1 - a).Add(f64.Vec3{0.5, 0.7, 1}.SMul(a))
	return rgbaFrom(v)
}

func Render(c Config) image.Image {
	imageHeight := max(1, int(float64(c.ImageWidth)/c.AspectRatio))

	world := Hittables{
		Sphere{f64.Vec3{0, 0, -1}, 0.5},
		Sphere{f64.Vec3{0, -100.5, -1}, 100},
	}

	viewportWidth := c.ViewportHeight * (float64(c.ImageWidth) / float64(imageHeight))

	viewportU := f64.Vec3{viewportWidth, 0, 0}
	viewportV := f64.Vec3{0, -c.ViewportHeight, 0}

	pixelDeltaU := viewportU.SDiv(float64(c.ImageWidth))
	pixelDeltaV := viewportV.SDiv(float64(imageHeight))

	// = CameraCenter - <0, 0, FocalLength> - ViewportU / 2 - ViewportV / 2
	viewportUpperLeft := c.CameraCenter.
		Sub(f64.Vec3{0, 0, c.FocalLength}).
		Sub(viewportU.SDiv(2)).
		Sub(viewportV.SDiv(2))

	// = ViewportUpperLeft + 0.5 * (PixelDeltaU + PixelDeltaV)
	pixel00Loc := viewportUpperLeft.Add(pixelDeltaU.Add(pixelDeltaV).SMul(0.5))

	img := image.NewRGBA(image.Rect(0, 0, c.ImageWidth, imageHeight))
	for j := 0; j < imageHeight; j++ {
		log.Printf("Scanlines remaining: %d", imageHeight-j)
		for i := 0; i < c.ImageWidth; i++ {
			// =  pixel00Loc + (i * pixelDeltaU) + (j * pixelDeltaV);
			pixelCenter := pixel00Loc.
				Add(pixelDeltaU.SMul(float64(i))).
				Add(pixelDeltaV.SMul(float64(j)))

			rayDirection := pixelCenter.Sub(c.CameraCenter)

			ray := Ray{c.CameraCenter, rayDirection}
			pixelColor := rayColor(world, ray)

			img.SetRGBA(i, j, pixelColor)
		}
	}
	return img
}
