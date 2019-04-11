// Copyright 2019 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

// Package encoder provides encoding functionality for Base10 to Base62
// conversion with/without paddings
//
package encoder

import (
	"fmt"
	"strconv"
)

const (
	maxBase62 = uint64(62)
	mapping   = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// ToBase62 converts int types to Base62 encoded string
func ToBase62(u uint64) string {
	var a [65]byte // 64 + 1: +1 for sign of 64bit value in base 2
	i := len(a)
	for u >= maxBase62 {
		i--
		// Avoid using r = a%b in addition to q = a/maxBase62
		// since 64bit division and modulo operations
		// are calculated by runtime functions on 32bit machines.
		q := u / maxBase62
		a[i] = mapping[u-q*maxBase62]
		u = q
	}
	// when u < maxBase62
	i--
	a[i] = mapping[u]
	return string(a[i:])
}

// ToBase62WithPaddingZeros converts int types to Base62 encoded string with
// padding zeros
func ToBase62WithPaddingZeros(u uint, padding int64) string {
	formatter := "%+0" + strconv.FormatInt(padding, 10) + "s"
	return fmt.Sprintf(formatter, ToBase62(uint64(u)))
}
