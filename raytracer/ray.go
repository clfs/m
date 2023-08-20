package raytracer

type Ray struct {
	Origin, Direction Vec3
}

func (r Ray) At(t float64) *Vec3 {
	var v Vec3
	// v = r.Origin + r.Direction * t
	return v.Add(&r.Origin, v.SMul(&r.Direction, t))
}
