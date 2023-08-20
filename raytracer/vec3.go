package raytracer

import "math"

type Vec3 [3]float64

// X returns the x-component.
func (v *Vec3) X() float64 {
	return v[0]
}

// Y returns the y-component.
func (v *Vec3) Y() float64 {
	return v[1]
}

// Z returns the z-component.
func (v *Vec3) Z() float64 {
	return v[2]
}

// Negate sets v to the negation of v0.
func (v *Vec3) Negate(v0 *Vec3) {
	v[0] = -v0[0]
	v[1] = -v0[1]
	v[2] = -v0[2]
}

// Add sets v to the sum of v0 and v1.
func (v *Vec3) Add(v0, v1 *Vec3) {
	v[0] = v0[0] + v1[0]
	v[1] = v0[1] + v1[1]
	v[2] = v0[2] + v1[2]
}

// SMul sets v to the scalar multiplication of v0 and s.
func (v *Vec3) SMul(v0 *Vec3, s float64) {
	v[0] = v0[0] * s
	v[1] = v0[1] * s
	v[2] = v0[2] * s
}

// SDiv sets v to the scalar division of v0 and s.
func (v *Vec3) SDiv(v0 *Vec3, s float64) {
	v[0] = v0[0] / s
	v[1] = v0[1] / s
	v[2] = v0[2] / s
}

// Mag returns the magnitude of v.
func (v *Vec3) Mag() float64 {
	return math.Sqrt(v.MagSq())
}

// MagSq returns the magnitude squared of v.
func (v *Vec3) MagSq() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// HProd sets v to the Hadamard product of v0 and v1.
func (v *Vec3) HProd(v0, v1 *Vec3) {
	v[0] = v0[0] * v1[0]
	v[1] = v0[1] * v1[1]
	v[2] = v0[2] * v1[2]
}

// Dot returns the dot product of v0 and v1.
func Dot(v0, v1 *Vec3) float64 {
	return v0[0]*v1[0] + v0[1]*v1[1] + v0[2]*v1[2]
}

// Cross sets v to the cross product of v0 and v1.
func (v *Vec3) Cross(v0, v1 *Vec3) {
	v[0] = v0[1]*v1[2] - v0[2]*v1[1]
	v[1] = v0[2]*v1[0] - v0[0]*v1[2]
	v[2] = v0[0]*v1[1] - v0[1]*v1[0]
}

// Normalize sets v to the unit vector in the direction of v.
func (v *Vec3) Normalize() {
	v.SDiv(v, v.Mag())
}

// Unit sets v to the unit vector of v0.
func (v *Vec3) Unit(v0 *Vec3) {
	v0.SDiv(v0, v0.Mag())
}
