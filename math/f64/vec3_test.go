package f64

import "testing"

func TestVec3_Neg(t *testing.T) {
	v := Vec3{1, 2, 3}
	want := Vec3{-1, -2, -3}
	if got := v.Neg(); got != want {
		t.Errorf("%v.Negation() = %v, want %v", v, got, want)
	}
}

func TestVec3_Add(t *testing.T) {
	v := Vec3{1, 2, 3}
	u := Vec3{4, 5, 6}
	want := Vec3{5, 7, 9}
	if got := v.Add(u); got != want {
		t.Errorf("%v.Add(%v) = %v, want %v", v, u, got, want)
	}
}

func TestVec3_Sub(t *testing.T) {
	v := Vec3{1, 2, 3}
	u := Vec3{4, 5, 6}
	want := Vec3{-3, -3, -3}
	if got := v.Sub(u); got != want {
		t.Errorf("%v.Sub(%v) = %v, want %v", v, u, got, want)
	}
}

func TestVec3_SMul(t *testing.T) {
	v := Vec3{1, 2, 3}
	s := 4.0
	want := Vec3{4, 8, 12}
	if got := v.SMul(s); got != want {
		t.Errorf("%v.SMul(%v) = %v, want %v", v, s, got, want)
	}
}

func TestVec3_SDiv(t *testing.T) {
	v := Vec3{4, 8, 12}
	s := 4.0
	want := Vec3{1, 2, 3}
	if got := v.SDiv(s); got != want {
		t.Errorf("%v.SDiv(%v) = %v, want %v", v, s, got, want)
	}
}

func TestVec3_Mag(t *testing.T) {
	v := Vec3{3, 4, 12}
	want := 13.0
	if got := v.Mag(); got != want {
		t.Errorf("%v.Mag() = %v, want %v", v, got, want)
	}
}

func TestVec3_MagSq(t *testing.T) {
	v := Vec3{3, 4, 12}
	want := 169.0
	if got := v.MagSq(); got != want {
		t.Errorf("%v.MagSq() = %v, want %v", v, got, want)
	}
}

func TestVec3_HProd(t *testing.T) {
	v := Vec3{1, 2, 3}
	u := Vec3{4, 5, 6}
	want := Vec3{4, 10, 18}
	if got := v.HProd(u); got != want {
		t.Errorf("%v.HProd(%v) = %v, want %v", v, u, got, want)
	}
}

func TestVec3_Dot(t *testing.T) {
	v := Vec3{1, 2, 3}
	u := Vec3{4, 5, 6}
	want := 32.0
	if got := v.Dot(u); got != want {
		t.Errorf("%v.Dot(%v) = %v, want %v", v, u, got, want)
	}
}

func TestVec3_Cross(t *testing.T) {
	v := Vec3{1, 2, 3}
	u := Vec3{4, 5, 6}
	want := Vec3{-3, 6, -3}
	if got := v.Cross(u); got != want {
		t.Errorf("%v.Cross(%v) = %v, want %v", v, u, got, want)
	}
}

func TestVec3_Unit(t *testing.T) {
	v := Vec3{3, 4, 12}
	want := Vec3{3.0 / 13.0, 4.0 / 13.0, 12 / 13.0}
	if got := v.Unit(); got != want {
		t.Errorf("%v.Unit() = %v, want %v", v, got, want)
	}
}
