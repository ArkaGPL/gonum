// Code generated by "go generate github.com/ArkaGPL/gonum/unit”; DO NOT EDIT.

// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"errors"
	"fmt"
	"math"
	"unicode/utf8"
)

// Pressure represents a pressure in Pascals.
type Pressure float64

const Pascal Pressure = 1

// Unit converts the Pressure to a *Unit.
func (pr Pressure) Unit() *Unit {
	return New(float64(pr), Dimensions{
		LengthDim: -1,
		MassDim:   1,
		TimeDim:   -2,
	})
}

// Pressure allows Pressure to implement a Pressurer interface.
func (pr Pressure) Pressure() Pressure {
	return pr
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension.
func (pr *Pressure) From(u Uniter) error {
	if !DimensionsMatch(u, Pascal) {
		*pr = Pressure(math.NaN())
		return errors.New("unit: dimension mismatch")
	}
	*pr = Pressure(u.Unit().Value())
	return nil
}

func (pr Pressure) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", pr, float64(pr))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " Pa"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(pr))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(pr))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(pr))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(pr))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g Pa)", c, pr, float64(pr))
	}
}
