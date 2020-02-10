// Copyright 2020 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"time"

	"github.com/mustafaturan/monoton/mtimer"
)

// NewSecond returns the preconfigured second sequencer
func NewSecond() *Sequence {
	second := uint64(time.Second)
	timer := mtimer.New()
	return &Sequence{
		now:     func() uint64 { return timer.Now() / second },
		max:     62*62*62*62*62*62 - 1,
		maxTime: 62*62*62*62*62*62 - 1,
	}
}
