// Copyright 2019 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"time"

	"github.com/mustafaturan/monoton/mtimer"
)

// NewMillisecond returns the preconfigured millisecond sequencer
func NewMillisecond() *Sequence {
	millisecond := uint(time.Millisecond)
	return &Sequence{
		now:     func() uint { return mtimer.Now() / millisecond },
		max:     62*62*62*62 - 1,
		maxTime: 62*62*62*62*62*62*62*62 - 1,
	}
}
