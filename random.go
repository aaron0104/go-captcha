// Copyright 2011 Dmitry Chestnykh. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package captcha

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

// RandomDigits returns a byte slice of the given length containing
// pseudorandom numbers in range 0-9. The slice can be used as a captcha
// solution.
func RandomDigits(length int) []byte {
	return randomBytesMod(length, 10)
}

func DigitString(digits []byte) (s string, err error) {
	buf := make([]byte, len(digits))
	for i, d := range digits {
		if d > 9 {
			err = errors.New(fmt.Sprintf("captcha: %d is invalid digit, can only use 0-9", d))
			return
		}
		buf[i] = '0' + d
	}
	s = string(buf)
	return
}

// randomBytes returns a byte slice of the given length read from CSPRNG.
func randomBytes(length int) (b []byte) {
	b = make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("captcha: error reading random source: " + err.Error())
	}
	return
}

// randomBytesMod returns a byte slice of the given length, where each byte is
// a random number modulo mod.
func randomBytesMod(length int, mod byte) (b []byte) {
	b = make([]byte, length)
	maxrb := byte(256 - (256 % int(mod)))
	i := 0
	for {
		r := randomBytes(length + (length / 4))
		for _, c := range r {
			if c >= maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = c % mod
			i++
			if i == length {
				return
			}
		}
	}
	panic("unreachable")
}

var prng = &siprng{}

// randIntn returns a pseudorandom non-negative int in range [0, n).
func randIntn(n int) int {
	return prng.Intn(n)
}

// randInt returns a pseudorandom int in range [from, to].
func randInt(from, to int) int {
	return prng.Intn(to+1-from) + from
}

// randFloat returns a pseudorandom float64 in range [from, to].
func randFloat(from, to float64) float64 {
	return (to-from)*prng.Float64() + from
}
