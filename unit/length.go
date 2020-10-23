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

// Length represents a length in metres.
type Length float64

const Metre Length = 1

// Unit converts the Length to a *Unit.
func (l Length) Unit() *Unit {
	return New(float64(l), Dimensions{
		LengthDim: 1,
	})
}

// Length allows Length to implement a Lengther interface.
func (l Length) Length() Length {
	return l
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension.
func (l *Length) From(u Uniter) error {
	if !DimensionsMatch(u, Metre) {
		*l = Length(math.NaN())
		return errors.New("unit: dimension mismatch")
	}
	*l = Length(u.Unit().Value())
	return nil
}

func (l Length) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", l, float64(l))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " m"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(l))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(l))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(l))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(l))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g m)", c, l, float64(l))
	}
}
