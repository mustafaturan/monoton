// Copyright 2019 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

// Package monoton is a highly scalable, single/multi node, human-readable,
// predictable and incremental unique id generator
//
package monoton

import (
	"fmt"
	"math"
	"sync"

	"github.com/mustafaturan/monoton/encoder"
	"github.com/mustafaturan/monoton/sequencer"
)

const (
	totalByteSize       = 16
	maxNodeErrorMsg     = "Node can't be greater than %d (given %d)"
	maxByteSizeErrorMsg = "Total byte size can't be >= to sum of s:%d, t:%d"
)

type config struct {
	sync.Mutex
	sequencer       sequencer.Sequencer
	node            string
	timeSeqByteSize int64
	seqByteSize     int64
}

var c config

func init() {
	Configure(sequencer.NewMillisecond(), 0)
}

// Configure configures the monoton with the given generator and node. If you
// need to reset the node, then you have to reconfigure. If you do not configure
// the node then the node will be set to zero value.
func Configure(s sequencer.Sequencer, node uint) error {
	c = config{sequencer: s}

	if err := configureByteSizes(); err != nil {
		return err
	}
	if err := configureNode(node); err != nil {
		return err
	}

	return nil
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
	t, seq := next()

	return encoder.ToBase62WithPaddingZeros(t, c.timeSeqByteSize) +
		encoder.ToBase62WithPaddingZeros(seq, c.seqByteSize) +
		c.node
}

func next() (uint, uint) {
	c.Lock()
	defer c.Unlock()

	return c.sequencer.Next()
}

func configureByteSizes() error {
	maxSeqTime := encoder.ToBase62(uint64(c.sequencer.MaxSequenceTime()))
	c.timeSeqByteSize = int64(len(maxSeqTime))

	maxSeq := encoder.ToBase62(uint64(c.sequencer.MaxSequence()))
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
