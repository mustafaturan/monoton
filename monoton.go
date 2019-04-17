// Copyright 2019 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

// Package monoton is a highly scalable, single/multi node, human-readable,
// predictable and incremental unique id generator
//
// Time Ordered
//
// The monoton package provides sequences based on the monotonic time which
// represents the absolute elapsed wall-clock time since some arbitrary, fixed
// point in the past. It isn't affected by changes in the system time-of-day
// clock.
//
// Initial Time
//
// Initial time value opens space for the time value by subtracting the given
// value from the time sequence.
//
// Readable
//
// The monoton package converts all sequences into Base62 format. And Base62
// only uses ASCII alpha-numeric chars to represent data which makes it easy to
// read, predict the order by a human eye.
//
// The total byte size is fixed to 16 bytes for all sequencers. And at least one
// byte is reserved to nodes.
//
// Multi Node Support
//
// The monoton package can be used on single/multiple nodes without the need for
// machine coordination. It uses configured node identifier to generate ids by
// attaching the node identifier to the end of the sequences.
//
// Extendable
//
// The package comes with three pre-configured sequencers and Sequencer
// interface to allow new sequencers.
//
// Included Sequencers and Byte Orderings
//
// The monoton package currently comes with Nanosecond, Millisecond and
// Second sequencers. And it uses Millisecond sequencer by default. For each
// sequencer, the byte orders are as following:
//
//	Second:      16 B =>  6 B (seconds)      + 6 B (counter) + 4 B (node)
//	Millisecond: 16 B =>  8 B (milliseconds) + 4 B (counter) + 4 B (node)
//	Nanosecond:  16 B => 11 B (nanoseconds)  + 2 B (counter) + 3 B (node)
//
// New Sequencers
//
// The sequencers can be extended for any other time format, sequence format by
// implementing the monoton/sequncer.Sequencer interface.
//
package monoton

import (
	"fmt"
	"math"

	"github.com/mustafaturan/monoton/encoder"
	"github.com/mustafaturan/monoton/sequencer"
)

const (
	totalByteSize       = 16
	maxNodeErrorMsg     = "node can't be greater than %d (given %d)"
	maxByteSizeErrorMsg = "sum of s:%d, t:%d bytes can't be >= total byte size"
)

type config struct {
	sequencer       *sequencer.Sequencer
	initialTime     uint
	node            string
	timeSeqByteSize int64
	seqByteSize     int64
}

var c config

// Configure configures the monoton with the given generator and node. If you
// need to reset the node, then you have to reconfigure. If you do not configure
// the node then the node will be set to zero value.
func Configure(s sequencer.Sequencer, node, initialTime uint) error {
	c = config{sequencer: &s, initialTime: initialTime}

	if err := configureByteSizes(); err != nil {
		return err
	}
	err := configureNode(node)

	return err
}

// Next generates next incremental unique identifier as Base62
// The execution will return the following Bytes (B) for the known sequencer
// types:
//
// 	Second:      16 B =>  6 B (seconds)      + 6 B (counter) + 4 B (node)
// 	Millisecond: 16 B =>  8 B (milliseconds) + 4 B (counter) + 4 B (node)
// 	Nanosecond:  16 B => 11 B (nanoseconds)  + 2 B (counter) + 3 B (node)
//
// For byte size decisions please refer to docs/adrs/byte-sizes.md
func Next() string {
	t, seq := (*c.sequencer).Next()

	return encoder.ToBase62WithPaddingZeros(t-c.initialTime, c.timeSeqByteSize) +
		encoder.ToBase62WithPaddingZeros(seq, c.seqByteSize) +
		c.node
}

func configureByteSizes() error {
	sequencer := *c.sequencer
	maxSeqTime := encoder.ToBase62(uint64(sequencer.MaxTime()))
	c.timeSeqByteSize = int64(len(maxSeqTime))

	maxSeq := encoder.ToBase62(uint64(sequencer.Max()))
	c.seqByteSize = int64(len(maxSeq))

	// At least one byte slot is necessary for the node
	if c.seqByteSize+c.timeSeqByteSize >= totalByteSize {
		return fmt.Errorf(maxByteSizeErrorMsg, c.seqByteSize, c.timeSeqByteSize)
	}

	return nil
}

func configureNode(node uint) error {
	nodeByteSize := totalByteSize - (c.timeSeqByteSize + c.seqByteSize)

	if err := validateNode(node, nodeByteSize); err != nil {
		return err
	}

	c.node = encoder.ToBase62WithPaddingZeros(node, nodeByteSize)
	return nil
}

func validateNode(node uint, nodeByteSize int64) error {
	maxNode := uint(math.Pow(62, float64(nodeByteSize))) - 1

	if node > maxNode {
		return fmt.Errorf(maxNodeErrorMsg, maxNode, node)
	}

	return nil
}
