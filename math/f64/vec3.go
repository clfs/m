package f64

import "math"

// Vec3 is a 3-element vector.
type Vec3 [3]float64

// Neg returns -v.
func (v Vec3) Neg() Vec3 {
	return Vec3{-v[0], -v[1], -v[2]}
}

// Add returns v + u.
func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{v[0] + u[0], v[1] + u[1], v[2] + u[2]}
}

// Sub returns v - u.
func (v Vec3) Sub(u Vec3) Vec3 {
	return Vec3{v[0] - u[0], v[1] - u[1], v[2] - u[2]}
}

// SMul returns v * s.
func (v Vec3) SMul(s float64) Vec3 {
	return Vec3{v[0] * s, v[1] * s, v[2] * s}
}

// SDiv returns v / s.
func (v Vec3) SDiv(s float64) Vec3 {
	return Vec3{v[0] / s, v[1] / s, v[2] / s}
}

// Mag returns the magnitude of v.
func (v Vec3) Mag() float64 {
	return math.Sqrt(v.MagSq())
}

// MagSq returns the magnitude squared of v.
func (v Vec3) MagSq() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// HProd returns the Hadamard product of v and u.
func (v Vec3) HProd(u Vec3) Vec3 {
	return Vec3{v[0] * u[0], v[1] * u[1], v[2] * u[2]}
}

// Dot returns the dot product of v and u.
func (v Vec3) Dot(u Vec3) float64 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2]
}

// Cross returns the cross product of v and u.
func (v Vec3) Cross(u Vec3) Vec3 {
	return Vec3{
		v[1]*u[2] - v[2]*u[1],
		v[2]*u[0] - v[0]*u[2],
		v[0]*u[1] - v[1]*u[0],
	}
}

// Unit returns the unit vector in the direction of of v.
func (v Vec3) Unit() Vec3 {
	return v.SDiv(v.Mag())
}
