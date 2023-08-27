package raytracer

import (
	"math"

	"github.com/clfs/m/math/f64"
)

type Sphere struct {
	Center f64.Vec3
	Radius float64
}

func (s Sphere) Hit(rec *HitRecord, ray Ray, tMin, tMax float64) bool {
	oc := ray.Origin.Sub(s.Center)
	a := ray.Direction.MagSq()
	halfB := oc.Dot(ray.Direction)
	c := oc.MagSq() - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false
	}
	sqrtD := math.Sqrt(discriminant)

	root := (-halfB - sqrtD) / a
	if root <= tMin || tMax <= root {
		root = (-halfB + sqrtD) / a
		if root <= tMin || tMax <= root {
			return false
		}
	}

	rec.T = root
	rec.Point = ray.At(rec.T)
	rec.Normal = rec.Point.Sub(s.Center).SDiv(s.Radius)

	return true
}
