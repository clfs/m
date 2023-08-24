package raytracer

import "github.com/clfs/m/math/f64"

type Ray struct {
	Origin    f64.Vec3
	Direction f64.Vec3
}

func (r Ray) At(t float64) f64.Vec3 {
	// r.Origin + r.Direction * t
	return r.Origin.Add(r.Direction.SMul(t))
}
