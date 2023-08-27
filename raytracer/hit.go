package raytracer

import "github.com/clfs/m/math/f64"

type Hittable interface {
	Hit(rec *HitRecord, ray Ray, tMin, tMax float64) bool
}

type HitRecord struct {
	Point  f64.Vec3
	Normal f64.Vec3
	T      float64
}
