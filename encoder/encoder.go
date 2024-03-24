// Copyright 2021 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

// Package encoder provides encoding functionality for Base10 to Base62
// conversion with/without paddings
package encoder

const (
	maxBase62 = uint64(62)
	mapping   = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// ToBase62WithPaddingZeros converts int types to Base62 encoded byte array
// with padding zeros
func ToBase62WithPaddingZeros(u uint64, length int) []byte {
	const size = 11 // largest uint64 in base62 occupies 11 bytes
	var a [size]byte
	i := size
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
	for i > size-length {
		i--
		a[i] = mapping[0]
	}
	return a[i:]
}

// Base62ByteSize returns the minimum byte size length requirement to allocate
// the given unsigned integer's value
func Base62ByteSize(u uint64) int {
	i := 0
	for u >= maxBase62 {
		i++
		q := u / maxBase62
		u = q
	}
	return i + 1
}
