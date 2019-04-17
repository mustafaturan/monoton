// Copyright 2019 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

// Package sequencer provides sequences based on monotonic time
//
// Time
//
// The sequncer package provides sequences based on the monotonic time which
// represents the absolute elapsed wall-clock time since some arbitrary, fixed
// point in the past. It isn't affected by changes in the system time-of-day
// clock.
//
// Time - Consequences
//
// Since the monotonic time needs extra calculation steps when it is compared
// to regular system time, it also consumes an extra time while generating
// sequences.
//
// NOTE: According to the documentation of Go language
// [time package](https://golang.org/pkg/time/), on some systems the monotonic
// clock will stop if the computer goes to sleep. On such a system, t.Sub(u)
// may not accurately reflect the actual time that passed between t and u which
// will result with incorrect sequences.
//
// Byte Sizes
//
// The total byte size is fixed to 16 bytes for any sequencer. And at least one
// byte is reserved to nodes. The package comes with three pre-configured
// sequencers and Sequencer interface to allow new sequencers.
//
// Byte Sizes - Defaults
//
// The package comes with pre-configured byte sizes for the Nanosecond,
// Millisecond and Second sequencers. And it does not allow you to adjust
// current sizes unless you create another sequencer.
// The defaults are adjusted the time and sequence byte sizes depending on
// general needs and to increase compatibility between projects.
//
// The current byte sizes:
//
// 	Second:      16 B =>  6 B (seconds)      + 6 B (counter) + 4 B (node)
// 	Millisecond: 16 B =>  8 B (milliseconds) + 4 B (counter) + 4 B (node)
// 	Nanosecond:  16 B => 11 B (nanoseconds)  + 2 B (counter) + 3 B (node)
//
// Byte Sizes - Consequences
//
// Although a strict byte size is limiting the space for nodes and sequences,
// 16 B gives enough flexibility for time, counter and nodes. In next 50 years,
// it could be necessary to provide a strategy to upgrade byte size to 32 B.
//
package sequencer

// Sequencer is a generic behavior for the sequence generators
type Sequencer interface {
	// Max returns the maximum possible sequence value for a given time
	Max() uint
	// MaxTime returns the maximum possible time sequence value
	MaxTime() uint
	// Now returns the current monotonic time
	Next() (uint, uint)
}
