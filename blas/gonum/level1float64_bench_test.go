// Code generated by "go run github.com/ArkaGPL/gonum/blas/testblas/benchautogen/autogen_bench_level1double.go"; DO NOT EDIT.

// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file

package gonum

import (
	"testing"

	"golang.org/x/exp/rand"

	"github.com/ArkaGPL/gonum/blas"
)

const (
	posInc1      = 5
	posInc2      = 3
	negInc1      = -3
	negInc2      = -4
	SMALL_SLICE  = 10
	MEDIUM_SLICE = 1000
	LARGE_SLICE  = 100000
	HUGE_SLICE   = 10000000
)

func randomSlice(l, idx int) []float64 {
	if idx < 0 {
		idx = -idx
	}
	s := make([]float64, l*idx)
	for i := range s {
		s[i] = rand.Float64()
	}
	return s
}

func benchmarkDdot(b *testing.B, n int, x []float64, incX int, y []float64, incY int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Ddot(n, x, incX, y, incY)
	}
}

func BenchmarkDdotSmallBothUnitary(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotSmallIncUni(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotSmallUniInc(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotSmallBothInc(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotMediumBothUnitary(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotMediumIncUni(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotMediumUniInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotMediumBothInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotLargeBothUnitary(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotLargeIncUni(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotLargeUniInc(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotLargeBothInc(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotHugeBothUnitary(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotHugeIncUni(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotHugeUniInc(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

func BenchmarkDdotHugeBothInc(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDdot(b, n, x, incX, y, incY)
}

/* ------------------ */
func benchmarkDnrm2(b *testing.B, n int, x []float64, incX int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Dnrm2(n, x, incX)
	}
}

func BenchmarkDnrm2SmallUnitaryInc(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)

	benchmarkDnrm2(b, n, x, incX)
}

func BenchmarkDnrm2SmallPosInc(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)

	benchmarkDnrm2(b, n, x, incX)
}

func BenchmarkDnrm2MediumUnitaryInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)

	benchmarkDnrm2(b, n, x, incX)
}

func BenchmarkDnrm2MediumPosInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)

	benchmarkDnrm2(b, n, x, incX)
}

func BenchmarkDnrm2LargeUnitaryInc(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)

	benchmarkDnrm2(b, n, x, incX)
}

func BenchmarkDnrm2LargePosInc(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)

	benchmarkDnrm2(b, n, x, incX)
}

func BenchmarkDnrm2HugeUnitaryInc(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)

	benchmarkDnrm2(b, n, x, incX)
}

func BenchmarkDnrm2HugePosInc(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)

	benchmarkDnrm2(b, n, x, incX)
}

/* ------------------ */
func benchmarkDasum(b *testing.B, n int, x []float64, incX int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Dasum(n, x, incX)
	}
}

func BenchmarkDasumSmallUnitaryInc(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)

	benchmarkDasum(b, n, x, incX)
}

func BenchmarkDasumSmallPosInc(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)

	benchmarkDasum(b, n, x, incX)
}

func BenchmarkDasumMediumUnitaryInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)

	benchmarkDasum(b, n, x, incX)
}

func BenchmarkDasumMediumPosInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)

	benchmarkDasum(b, n, x, incX)
}

func BenchmarkDasumLargeUnitaryInc(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)

	benchmarkDasum(b, n, x, incX)
}

func BenchmarkDasumLargePosInc(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)

	benchmarkDasum(b, n, x, incX)
}

func BenchmarkDasumHugeUnitaryInc(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)

	benchmarkDasum(b, n, x, incX)
}

func BenchmarkDasumHugePosInc(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)

	benchmarkDasum(b, n, x, incX)
}

/* ------------------ */
func benchmarkIdamax(b *testing.B, n int, x []float64, incX int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Idamax(n, x, incX)
	}
}

func BenchmarkIdamaxSmallUnitaryInc(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)

	benchmarkIdamax(b, n, x, incX)
}

func BenchmarkIdamaxSmallPosInc(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)

	benchmarkIdamax(b, n, x, incX)
}

func BenchmarkIdamaxMediumUnitaryInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)

	benchmarkIdamax(b, n, x, incX)
}

func BenchmarkIdamaxMediumPosInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)

	benchmarkIdamax(b, n, x, incX)
}

func BenchmarkIdamaxLargeUnitaryInc(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)

	benchmarkIdamax(b, n, x, incX)
}

func BenchmarkIdamaxLargePosInc(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)

	benchmarkIdamax(b, n, x, incX)
}

func BenchmarkIdamaxHugeUnitaryInc(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)

	benchmarkIdamax(b, n, x, incX)
}

func BenchmarkIdamaxHugePosInc(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)

	benchmarkIdamax(b, n, x, incX)
}

/* ------------------ */
func benchmarkDswap(b *testing.B, n int, x []float64, incX int, y []float64, incY int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Dswap(n, x, incX, y, incY)
	}
}

func BenchmarkDswapSmallBothUnitary(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapSmallIncUni(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapSmallUniInc(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapSmallBothInc(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapMediumBothUnitary(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapMediumIncUni(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapMediumUniInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapMediumBothInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapLargeBothUnitary(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapLargeIncUni(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapLargeUniInc(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapLargeBothInc(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapHugeBothUnitary(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapHugeIncUni(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapHugeUniInc(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

func BenchmarkDswapHugeBothInc(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDswap(b, n, x, incX, y, incY)
}

/* ------------------ */
func benchmarkDcopy(b *testing.B, n int, x []float64, incX int, y []float64, incY int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Dcopy(n, x, incX, y, incY)
	}
}

func BenchmarkDcopySmallBothUnitary(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopySmallIncUni(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopySmallUniInc(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopySmallBothInc(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopyMediumBothUnitary(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopyMediumIncUni(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopyMediumUniInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopyMediumBothInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopyLargeBothUnitary(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopyLargeIncUni(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopyLargeUniInc(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopyLargeBothInc(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopyHugeBothUnitary(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopyHugeIncUni(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopyHugeUniInc(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

func BenchmarkDcopyHugeBothInc(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)

	benchmarkDcopy(b, n, x, incX, y, incY)
}

/* ------------------ */
func benchmarkDaxpy(b *testing.B, n int, alpha float64, x []float64, incX int, y []float64, incY int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Daxpy(n, alpha, x, incX, y, incY)
	}
}

func BenchmarkDaxpySmallBothUnitary(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpySmallIncUni(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpySmallUniInc(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpySmallBothInc(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpyMediumBothUnitary(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpyMediumIncUni(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpyMediumUniInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpyMediumBothInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpyLargeBothUnitary(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpyLargeIncUni(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpyLargeUniInc(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpyLargeBothInc(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpyHugeBothUnitary(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpyHugeIncUni(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpyHugeUniInc(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

func BenchmarkDaxpyHugeBothInc(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	alpha := 2.4
	benchmarkDaxpy(b, n, alpha, x, incX, y, incY)
}

/* ------------------ */
func benchmarkDrot(b *testing.B, n int, x []float64, incX int, y []float64, incY int, c, s float64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Drot(n, x, incX, y, incY, c, s)
	}
}

func BenchmarkDrotSmallBothUnitary(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotSmallIncUni(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotSmallUniInc(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotSmallBothInc(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotMediumBothUnitary(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotMediumIncUni(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotMediumUniInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotMediumBothInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotLargeBothUnitary(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotLargeIncUni(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotLargeUniInc(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotLargeBothInc(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotHugeBothUnitary(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotHugeIncUni(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotHugeUniInc(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

func BenchmarkDrotHugeBothInc(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	c := 0.89725836967
	s := 0.44150585279
	benchmarkDrot(b, n, x, incX, y, incY, c, s)
}

/* ------------------ */
func benchmarkDrotmOffDia(b *testing.B, n int, x []float64, incX int, y []float64, incY int, p blas.DrotmParams) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Drotm(n, x, incX, y, incY, p)
	}
}

func BenchmarkDrotmOffDiaSmallBothUnitary(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaSmallIncUni(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaSmallUniInc(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaSmallBothInc(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaMediumBothUnitary(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaMediumIncUni(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaMediumUniInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaMediumBothInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaLargeBothUnitary(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaLargeIncUni(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaLargeUniInc(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaLargeBothInc(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaHugeBothUnitary(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaHugeIncUni(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaHugeUniInc(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmOffDiaHugeBothInc(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{0, -0.625, 0.9375, 0}}
	benchmarkDrotmOffDia(b, n, x, incX, y, incY, p)
}

/* ------------------ */
func benchmarkDrotmDia(b *testing.B, n int, x []float64, incX int, y []float64, incY int, p blas.DrotmParams) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Drotm(n, x, incX, y, incY, p)
	}
}

func BenchmarkDrotmDiaSmallBothUnitary(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaSmallIncUni(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaSmallUniInc(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaSmallBothInc(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaMediumBothUnitary(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaMediumIncUni(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaMediumUniInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaMediumBothInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaLargeBothUnitary(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaLargeIncUni(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaLargeUniInc(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaLargeBothInc(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaHugeBothUnitary(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaHugeIncUni(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaHugeUniInc(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmDiaHugeBothInc(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{5.0 / 12, 0, 0, 0.625}}
	benchmarkDrotmDia(b, n, x, incX, y, incY, p)
}

/* ------------------ */
func benchmarkDrotmResc(b *testing.B, n int, x []float64, incX int, y []float64, incY int, p blas.DrotmParams) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Drotm(n, x, incX, y, incY, p)
	}
}

func BenchmarkDrotmRescSmallBothUnitary(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescSmallIncUni(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescSmallUniInc(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescSmallBothInc(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescMediumBothUnitary(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescMediumIncUni(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescMediumUniInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescMediumBothInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescLargeBothUnitary(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescLargeIncUni(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescLargeUniInc(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescLargeBothInc(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescHugeBothUnitary(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescHugeIncUni(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := 1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescHugeUniInc(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

func BenchmarkDrotmRescHugeBothInc(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	incY := negInc1
	y := randomSlice(n, incY)
	p := blas.DrotmParams{Flag: blas.OffDiagonal, H: [4]float64{4096, -3584, 1792, 4096}}
	benchmarkDrotmResc(b, n, x, incX, y, incY, p)
}

/* ------------------ */
func benchmarkDscal(b *testing.B, n int, alpha float64, x []float64, incX int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Dscal(n, alpha, x, incX)
	}
}

func BenchmarkDscalSmallUnitaryInc(b *testing.B) {
	n := SMALL_SLICE
	incX := 1
	x := randomSlice(n, incX)
	alpha := 2.4
	benchmarkDscal(b, n, alpha, x, incX)
}

func BenchmarkDscalSmallPosInc(b *testing.B) {
	n := SMALL_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	alpha := 2.4
	benchmarkDscal(b, n, alpha, x, incX)
}

func BenchmarkDscalMediumUnitaryInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := 1
	x := randomSlice(n, incX)
	alpha := 2.4
	benchmarkDscal(b, n, alpha, x, incX)
}

func BenchmarkDscalMediumPosInc(b *testing.B) {
	n := MEDIUM_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	alpha := 2.4
	benchmarkDscal(b, n, alpha, x, incX)
}

func BenchmarkDscalLargeUnitaryInc(b *testing.B) {
	n := LARGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	alpha := 2.4
	benchmarkDscal(b, n, alpha, x, incX)
}

func BenchmarkDscalLargePosInc(b *testing.B) {
	n := LARGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	alpha := 2.4
	benchmarkDscal(b, n, alpha, x, incX)
}

func BenchmarkDscalHugeUnitaryInc(b *testing.B) {
	n := HUGE_SLICE
	incX := 1
	x := randomSlice(n, incX)
	alpha := 2.4
	benchmarkDscal(b, n, alpha, x, incX)
}

func BenchmarkDscalHugePosInc(b *testing.B) {
	n := HUGE_SLICE
	incX := posInc1
	x := randomSlice(n, incX)
	alpha := 2.4
	benchmarkDscal(b, n, alpha, x, incX)
}

/* ------------------ */
