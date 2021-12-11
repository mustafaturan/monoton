// Copyright 2021 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"github.com/mustafaturan/monoton/v3/mtimer"
)

// NewNanosecond returns the preconfigured nanosecond sequencer
func NewNanosecond() *Sequence {
	timer := mtimer.New()
	return &Sequence{
		now:     timer.Now,
		max:     62*62 - 1,
		maxTime: uint64(1<<64 - 1),
		maxNode: 62*62*62 - 1,
	}
}
