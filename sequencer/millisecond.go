// Copyright 2020 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"time"

	"github.com/mustafaturan/monoton/mtimer"
)

// NewMillisecond returns the preconfigured millisecond sequencer
func NewMillisecond() *Sequence {
	millisecond := uint64(time.Millisecond)
	timer := mtimer.New()
	return &Sequence{
		now:     func() uint64 { return timer.Now() / millisecond },
		max:     62*62*62*62 - 1,
		maxTime: 62*62*62*62*62*62*62*62 - 1,
	}
}
