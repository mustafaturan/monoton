// Copyright 2019 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"github.com/mustafaturan/monoton/mtimer"
)

// NewNanosecond returns the preconfigured nanosecond sequencer
func NewNanosecond() *Sequence {
	return &Sequence{
		now:     mtimer.Now,
		max:     62*62 - 1,
		maxTime: uint(1<<64 - 1),
	}
}
