package raytracer

import "testing"

func TestRay_At(t *testing.T) {
	cases := []struct {
		in   Ray
		t    float64
		want Vec3
	}{
		{
			Ray{
				Origin:    Vec3{0, 0, 0},
				Direction: Vec3{1, 0, 0},
			},
			1,
			Vec3{1, 0, 0},
		},
		{
			Ray{
				Origin:    Vec3{0, 0, 0},
				Direction: Vec3{1, 1, 1},
			},
			1,
			Vec3{1, 1, 1},
		},
		{
			Ray{
				Origin:    Vec3{0, 0, 0},
				Direction: Vec3{1, 1, 1},
			},
			2,
			Vec3{2, 2, 2},
		},
	}

	for i, tc := range cases {
		got := *tc.in.At(tc.t)
		if tc.want != got {
			t.Errorf("#%d: want %v, got %v", i, tc.want, got)
		}
	}
}
