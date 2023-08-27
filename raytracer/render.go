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

func hitSphere(center f64.Vec3, radius float64, r Ray) float64 {
	oc := r.Origin.Sub(center)
	a := r.Direction.Dot(r.Direction)
	b := 2 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1
	}
	return (-b - math.Sqrt(discriminant)) / (2 * a)
}

func rayColor(r Ray) color.RGBA {
	t := hitSphere(f64.Vec3{0, 0, -1}, 0.5, r)
	if t > 0 {
		unitNormal := r.At(t).Sub(f64.Vec3{0, 0, -1}).Unit()
		// = 0.5 * (<X, Y, Z> + <1, 1, 1>)
		return rgbaFrom(unitNormal.Add(f64.Vec3{1, 1, 1}).SMul(0.5))
	}

	unitDirection := r.Direction.Unit()
	a := 0.5 * (unitDirection[1] + 1)

	// = (1 - a) * <1, 1, 1> + a * <0.5, 0.7, 1>
	v := f64.Vec3{1, 1, 1}.SMul(1 - a).Add(f64.Vec3{0.5, 0.7, 1}.SMul(a))
	return rgbaFrom(v)
}

func Render(c Config) image.Image {
	imageHeight := max(1, int(float64(c.ImageWidth)/c.AspectRatio))

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
			pixelColor := rayColor(ray)

			img.SetRGBA(i, j, pixelColor)
		}
	}
	return img
}
