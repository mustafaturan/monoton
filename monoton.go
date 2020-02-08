// Copyright 2020 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

/*
Package monoton is a highly scalable, single/multi node, human-readable,
predictable and incremental unique id generator

Time Ordered

The monoton package provides sequences based on the monotonic time which
represents the absolute elapsed wall-clock time since some arbitrary, fixed
point in the past. It isn't affected by changes in the system time-of-day
clock.

Initial Time

Initial time value opens space for the time value by subtracting the given
value from the time sequence.

Readable

The monoton package converts all sequences into Base62 format. And Base62
only uses ASCII alpha-numeric chars to represent data which makes it easy to
read, predict the order by a human eye.

The total byte size is fixed to 16 bytes for all sequencers. And at least one
byte is reserved to nodes.

Multi Node Support

The monoton package can be used on single/multiple nodes without the need for
machine coordination. It uses configured node identifier to generate ids by
attaching the node identifier to the end of the sequences.

Extendable

The package comes with three pre-configured sequencers and Sequencer
interface to allow new sequencers.

Included Sequencers and Byte Orderings

The monoton package currently comes with Nanosecond, Millisecond and
Second sequencers. And it uses Millisecond sequencer by default. For each
sequencer, the byte orders are as following:

	Second:      16 B =>  6 B (seconds)      + 6 B (counter) + 4 B (node)
	Millisecond: 16 B =>  8 B (milliseconds) + 4 B (counter) + 4 B (node)
	Nanosecond:  16 B => 11 B (nanoseconds)  + 2 B (counter) + 3 B (node)

New Sequencers

The sequencers can be extended for any other time format, sequence format by
implementing the monoton/sequncer.Sequencer interface.

Example using Singleton

	package uniqid

	// Import packages
	import (
		"fmt"
		"github.com/mustafaturan/monoton"
		"github.com/mustafaturan/monoton/sequencer"
	)

	const year2020asMillisecondPST = 1577865600000

	var m monoton.Monoton

	// On init configure the monoton
	func init() {
		m = *(newIDGenerator()) // to store in the stack
	}

	func newIDGenerator() *monoton.Monoton {
		// Fetch your node id from a config server or generate from MAC/IP
		// address
		node := uint64(1)

		// A unix time value which will be subtracted from the time sequence
		// value. The initialTime value type corresponds to the sequencer type's
		// time representation. If you are using Millisecond sequencer then it
		// must be considered as Millisecond
		initialTime := uint64(year2020asMillisecondPST)

		// Configure monoton with a sequencer and the node
		m, err = monoton.New(sequencer.NewMillisecond(), node, initialTime)
		if err != nil{
			panic(err)
		}

		return m
	}

	func Generate() string {
		m.Next()
	}

	// In any other package unique ids can be generated like below:

	package main

	import (
		"fmt"
		"uniqid" // your local `uniqid` package inside your project
	)

	func main() {
		for i := 0; i < 100; i++ {
			fmt.Println(uniqid.Generate())
		}
	}

*/
package monoton

import (
	"fmt"
	"math"

	"github.com/mustafaturan/monoton/encoder"
	"github.com/mustafaturan/monoton/sequencer"
)

const (
	totalByteSize  = 16
	errMaxNode     = "node can't be greater than %d (given %d)"
	errMaxByteSize = "max byte size sum of sequence(%d) and time sequence(%d) can't be >= total byte size(%d)"
)

// MaxNodeCapacityExceededError is an error type with node information
type MaxNodeCapacityExceededError struct {
	Node    uint64
	MaxNode uint64
}

func (e *MaxNodeCapacityExceededError) Error() string {
	return fmt.Sprintf(errMaxNode, e.MaxNode, e.Node)
}

// MaxByteSizeError is an error type with sequence & time byte sizes
type MaxByteSizeError struct {
	ByteSizeSequence     int64
	ByteSizeSequenceTime int64
	ByteSizeTotal        int64

	MaxSequence     string
	MaxSequenceTime string
}

func (e *MaxByteSizeError) Error() string {
	return fmt.Sprintf(
		errMaxByteSize,
		e.ByteSizeSequence,
		e.ByteSizeSequenceTime,
		e.ByteSizeTotal,
	)
}

// Monoton is a sequential id generator
type Monoton struct {
	sequencer       *sequencer.Sequencer
	initialTime     uint64
	node            string
	timeSeqByteSize int64
	seqByteSize     int64
}

// New inits a new monoton ID generator with the given generator and node.
func New(s sequencer.Sequencer, node, initialTime uint64) (*Monoton, error) {
	m := &Monoton{sequencer: &s, initialTime: initialTime}

	if err := m.configureByteSizes(); err != nil {
		return nil, err
	}

	if err := m.configureNode(node); err != nil {
		return nil, err
	}

	return m, nil
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
func (m *Monoton) Next() string {
	t, seq := (*m.sequencer).Next()

	return encoder.ToBase62WithPaddingZeros(t-m.initialTime, m.timeSeqByteSize) +
		encoder.ToBase62WithPaddingZeros(seq, m.seqByteSize) +
		m.node
}

func (m *Monoton) configureByteSizes() error {
	sequencer := *m.sequencer
	maxSeqTime := encoder.ToBase62(sequencer.MaxTime())
	m.timeSeqByteSize = int64(len(maxSeqTime))

	maxSeq := encoder.ToBase62(sequencer.Max())
	m.seqByteSize = int64(len(maxSeq))

	// At least one byte slot is necessary for the node
	if m.seqByteSize+m.timeSeqByteSize >= totalByteSize {
		return &MaxByteSizeError{
			ByteSizeSequence:     m.seqByteSize,
			ByteSizeSequenceTime: m.timeSeqByteSize,
			ByteSizeTotal:        totalByteSize,
			MaxSequence:          maxSeq,
			MaxSequenceTime:      maxSeqTime,
		}
	}

	return nil
}

func (m *Monoton) configureNode(node uint64) error {
	if err := m.validateNode(node); err != nil {
		return err
	}

	m.node = encoder.ToBase62WithPaddingZeros(node, m.nodeByteSize())
	return nil
}

func (m *Monoton) validateNode(node uint64) error {
	maxNode := uint64(math.Pow(62, float64(m.nodeByteSize()))) - 1
	if node > maxNode {
		return &MaxNodeCapacityExceededError{Node: node, MaxNode: maxNode}
	}

	return nil
}

func (m *Monoton) nodeByteSize() int64 {
	return totalByteSize - (m.timeSeqByteSize + m.seqByteSize)
}
