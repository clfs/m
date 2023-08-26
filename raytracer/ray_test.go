package raytracer

import (
	"testing"

	"github.com/clfs/m/math/f64"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestRay_At(t *testing.T) {
	cases := []struct {
		r    Ray
		tt   float64
		want f64.Vec3
	}{
		{Ray{f64.Vec3{0, 0, 0}, f64.Vec3{1, 0, 0}}, 0, f64.Vec3{0, 0, 0}},
		{Ray{f64.Vec3{1, 2, 3}, f64.Vec3{-4, 5, 8}}, 3, f64.Vec3{-11, 17, 27}},
	}

	for i, tc := range cases {
		got := tc.r.At(tc.tt)
		diff := cmp.Diff(tc.want, got, cmpopts.EquateApprox(0, 0.0001))
		if diff != "" {
			t.Errorf("#%d: mismatch: (-want +got):\n%s", i, diff)
		}
	}
}
