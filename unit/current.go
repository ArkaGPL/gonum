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

// Current represents a current in Amperes.
type Current float64

const Ampere Current = 1

// Unit converts the Current to a *Unit.
func (i Current) Unit() *Unit {
	return New(float64(i), Dimensions{
		CurrentDim: 1,
	})
}

// Current allows Current to implement a Currenter interface.
func (i Current) Current() Current {
	return i
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension.
func (i *Current) From(u Uniter) error {
	if !DimensionsMatch(u, Ampere) {
		*i = Current(math.NaN())
		return errors.New("unit: dimension mismatch")
	}
	*i = Current(u.Unit().Value())
	return nil
}

func (i Current) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", i, float64(i))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " A"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(i))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(i))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(i))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(i))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g A)", c, i, float64(i))
	}
}
