package raytracer

import "github.com/clfs/m/math/f64"

type Hittable interface {
	Hit(rec *HitRecord, ray Ray, tMin, tMax float64) bool
}

// Hittables is a slice of Hittable that also implements Hittable.
type Hittables []Hittable

func (hs Hittables) Hit(rec *HitRecord, ray Ray, tMin, tMax float64) bool {
	var tempRec HitRecord
	var hitAnything bool
	closestSoFar := tMax

	for _, h := range hs {
		if h.Hit(&tempRec, ray, tMin, closestSoFar) {
			hitAnything = true
			closestSoFar = tempRec.T
			*rec = tempRec
		}
	}

	return hitAnything
}

type HitRecord struct {
	Point     f64.Vec3
	Normal    f64.Vec3
	T         float64
	FrontFace bool
}

// SetFaceNormal sets the hit record normal vector.
// outwardNormal must have unit length.
func (rec *HitRecord) SetFaceNormal(ray Ray, outwardNormal f64.Vec3) {
	rec.FrontFace = ray.Direction.Dot(outwardNormal) < 0
	if rec.FrontFace {
		rec.Normal = outwardNormal
	} else {
		rec.Normal = outwardNormal.Neg()
	}
}
