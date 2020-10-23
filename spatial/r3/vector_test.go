// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package r3

import (
	"math"
	"testing"

	"github.com/ArkaGPL/gonum/floats/scalar"
)

func TestAdd(t *testing.T) {
	for _, test := range []struct {
		v1, v2 Vec
		want   Vec
	}{
		{Vec{0, 0, 0}, Vec{0, 0, 0}, Vec{0, 0, 0}},
		{Vec{1, 0, 0}, Vec{0, 0, 0}, Vec{1, 0, 0}},
		{Vec{1, 2, 3}, Vec{4, 5, 7}, Vec{5, 7, 10}},
		{Vec{1, -3, 5}, Vec{1, -6, -6}, Vec{2, -9, -1}},
		{Vec{1, 2, 3}, Vec{-1, -2, -3}, Vec{}},
	} {
		t.Run("", func(t *testing.T) {
			got := test.v1.Add(test.v2)

			if got != test.want {
				t.Fatalf(
					"error: %v + %v: got=%v, want=%v",
					test.v1, test.v2, got, test.want,
				)
			}
		})
	}
}

func TestSub(t *testing.T) {
	for _, test := range []struct {
		v1, v2 Vec
		want   Vec
	}{
		{Vec{0, 0, 0}, Vec{0, 0, 0}, Vec{0, 0, 0}},
		{Vec{1, 0, 0}, Vec{0, 0, 0}, Vec{1, 0, 0}},
		{Vec{1, 2, 3}, Vec{4, 5, 7}, Vec{-3, -3, -4}},
		{Vec{1, -3, 5}, Vec{1, -6, -6}, Vec{0, 3, 11}},
		{Vec{1, 2, 3}, Vec{1, 2, 3}, Vec{}},
	} {
		t.Run("", func(t *testing.T) {
			got := test.v1.Sub(test.v2)
			if got != test.want {
				t.Fatalf(
					"error: %v - %v: got=%v, want=%v",
					test.v1, test.v2, got, test.want,
				)
			}
		})
	}
}

func TestScale(t *testing.T) {
	for _, test := range []struct {
		a    float64
		v    Vec
		want Vec
	}{
		{3, Vec{0, 0, 0}, Vec{0, 0, 0}},
		{1, Vec{1, 0, 0}, Vec{1, 0, 0}},
		{0, Vec{1, 0, 0}, Vec{0, 0, 0}},
		{3, Vec{1, 0, 0}, Vec{3, 0, 0}},
		{-1, Vec{1, -3, 5}, Vec{-1, 3, -5}},
		{2, Vec{1, -3, 5}, Vec{2, -6, 10}},
		{10, Vec{1, 2, 3}, Vec{10, 20, 30}},
	} {
		t.Run("", func(t *testing.T) {
			got := test.v.Scale(test.a)
			if got != test.want {
				t.Fatalf(
					"error: %v * %v: got=%v, want=%v",
					test.a, test.v, got, test.want)
			}
		})
	}
}

func TestDot(t *testing.T) {
	for _, test := range []struct {
		u, v Vec
		want float64
	}{
		{Vec{1, 2, 3}, Vec{1, 2, 3}, 14},
		{Vec{1, 0, 0}, Vec{1, 0, 0}, 1},
		{Vec{1, 0, 0}, Vec{0, 1, 0}, 0},
		{Vec{1, 0, 0}, Vec{0, 1, 1}, 0},
		{Vec{1, 1, 1}, Vec{-1, -1, -1}, -3},
		{Vec{1, 2, 2}, Vec{-0.3, 0.4, -1.2}, -1.9},
	} {
		t.Run("", func(t *testing.T) {
			{
				got := test.u.Dot(test.v)
				if got != test.want {
					t.Fatalf(
						"error: %v · %v: got=%v, want=%v",
						test.u, test.v, got, test.want,
					)
				}
			}
			{
				got := test.v.Dot(test.u)
				if got != test.want {
					t.Fatalf(
						"error: %v · %v: got=%v, want=%v",
						test.v, test.u, got, test.want,
					)
				}
			}
		})
	}
}

func TestCross(t *testing.T) {
	for _, test := range []struct {
		v1, v2, want Vec
	}{
		{Vec{1, 0, 0}, Vec{1, 0, 0}, Vec{0, 0, 0}},
		{Vec{1, 0, 0}, Vec{0, 1, 0}, Vec{0, 0, 1}},
		{Vec{0, 1, 0}, Vec{1, 0, 0}, Vec{0, 0, -1}},
		{Vec{1, 2, 3}, Vec{-4, 5, -6}, Vec{-27, -6, 13}},
		{Vec{1, 2, 3}, Vec{1, 2, 3}, Vec{}},
		{Vec{1, 2, 3}, Vec{2, 3, 4}, Vec{-1, 2, -1}},
	} {
		t.Run("", func(t *testing.T) {
			got := test.v1.Cross(test.v2)
			if got != test.want {
				t.Fatalf(
					"error: %v × %v = %v, want %v",
					test.v1, test.v2, got, test.want,
				)
			}
		})
	}
}

func TestNorm(t *testing.T) {
	for _, test := range []struct {
		v    Vec
		want float64
	}{
		{Vec{0, 0, 0}, 0},
		{Vec{0, 1, 0}, 1},
		{Vec{3, -4, 12}, 13},
		{Vec{1, 1e-16, 1e-32}, 1},
		{Vec{-0, 4.3145006366056343748277397783556100978621924913975e-196, 4.3145006366056343748277397783556100978621924913975e-196}, 6.101625315155041e-196},
	} {
		t.Run("", func(t *testing.T) {
			if got, want := Norm(test.v), test.want; got != want {
				t.Fatalf("|%v| = %v, want %v", test.v, got, want)
			}
		})
	}
}

func TestNorm2(t *testing.T) {
	for _, test := range []struct {
		v    Vec
		want float64
	}{
		{Vec{0, 0, 0}, 0},
		{Vec{0, 1, 0}, 1},
		{Vec{1, 1, 1}, 3},
		{Vec{1, 2, 3}, 14},
		{Vec{3, -4, 12}, 169},
		{Vec{1, 1e-16, 1e-32}, 1},
		// This will underflow and return zero.
		{Vec{-0, 4.3145006366056343748277397783556100978621924913975e-196, 4.3145006366056343748277397783556100978621924913975e-196}, 0},
	} {
		t.Run("", func(t *testing.T) {
			if got, want := Norm2(test.v), test.want; got != want {
				t.Fatalf("|%v|^2 = %v, want %v", test.v, got, want)
			}
		})
	}
}

func TestUnit(t *testing.T) {
	for _, test := range []struct {
		v, want Vec
	}{
		{Vec{}, Vec{math.NaN(), math.NaN(), math.NaN()}},
		{Vec{1, 0, 0}, Vec{1, 0, 0}},
		{Vec{0, 1, 0}, Vec{0, 1, 0}},
		{Vec{0, 0, 1}, Vec{0, 0, 1}},
		{Vec{1, 1, 1}, Vec{1. / math.Sqrt(3), 1. / math.Sqrt(3), 1. / math.Sqrt(3)}},
		{Vec{1, 1e-16, 1e-32}, Vec{1, 1e-16, 1e-32}},
	} {
		t.Run("", func(t *testing.T) {
			got := Unit(test.v)
			if !vecEqual(got, test.want) {
				t.Fatalf(
					"Normalize(%v) = %v, want %v",
					test.v, got, test.want,
				)
			}
			if test.v == (Vec{}) {
				return
			}
			if n, want := Norm(got), 1.0; n != want {
				t.Fatalf("|%v| = %v, want 1", got, n)
			}
		})
	}
}

func TestCos(t *testing.T) {
	for _, test := range []struct {
		v1, v2 Vec
		want   float64
	}{
		{Vec{1, 1, 1}, Vec{1, 1, 1}, 1},
		{Vec{1, 1, 1}, Vec{-1, -1, -1}, -1},
		{Vec{1, 1, 1}, Vec{1, -1, 1}, 1.0 / 3},
		{Vec{1, 0, 0}, Vec{1, 0, 0}, 1},
		{Vec{1, 0, 0}, Vec{0, 1, 0}, 0},
		{Vec{1, 0, 0}, Vec{0, 1, 1}, 0},
		{Vec{1, 0, 0}, Vec{-1, 0, 0}, -1},
	} {
		t.Run("", func(t *testing.T) {
			tol := 1e-14
			got := Cos(test.v1, test.v2)
			if !scalar.EqualWithinAbs(got, test.want, tol) {
				t.Fatalf("cos(%v, %v)= %v, want %v",
					test.v1, test.v2, got, test.want,
				)
			}
		})
	}
}

func vecIsNaN(v Vec) bool {
	return math.IsNaN(v.X) && math.IsNaN(v.Y) && math.IsNaN(v.Z)
}

func vecIsNaNAny(v Vec) bool {
	return math.IsNaN(v.X) || math.IsNaN(v.Y) || math.IsNaN(v.Z)
}

func vecEqual(a, b Vec) bool {
	if vecIsNaNAny(a) || vecIsNaNAny(b) {
		return vecIsNaN(a) && vecIsNaN(b)
	}
	return a == b
}
