# Monoton

[![Build Status](https://travis-ci.org/mustafaturan/monoton.svg?branch=master)](https://travis-ci.org/mustafaturan/monoton)
[![Coverage Status](https://coveralls.io/repos/github/mustafaturan/monoton/badge.svg?branch=master)](https://coveralls.io/github/mustafaturan/monoton?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/mustafaturan/monoton)](https://goreportcard.com/report/github.com/mustafaturan/monoton)
[![GoDoc](https://godoc.org/github.com/mustafaturan/monoton?status.svg)](https://godoc.org/github.com/mustafaturan/monoton/v3)

Highly scalable, single/multi node, predictable and incremental unique id
generator with zero allocation magic.

## Installation

Via go packages:
```go get github.com/mustafaturan/monoton/v3```

## API

The method names and arities/args are stable now. No change should be expected
on the package for the version `3.x.x` except any bug fixes.

## Usage

### Using with Singleton

Create a new package like below, and then call `Next()` or `NextBytes()` method:

```go
package uniqid

// Import packages
import (
	"fmt"
	"github.com/mustafaturan/monoton/v3"
	"github.com/mustafaturan/monoton/v3/sequencer"
)

var m monoton.Monoton

// On init configure the monoton
func init() {
	m = newIDGenerator()
}

func newIDGenerator() monoton.Monoton {
	// Fetch your node id from a config server or generate from MAC/IP address
	node := uint64(1)

	// A unix time value which will be subtracted from the time sequence value.
	// The initialTime value type corresponds to the sequencer type's time
	// representation. If you are using Millisecond sequencer then it must be
	// considered as Millisecond
	// If we want to init the time with 2020-01-01 00:00:00 PST
	initialTime := uint64(1577865600000)

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

func GeneateBytes() [16]byte {
	m.NextBytes()
}
```

In any other package generate the ids like below:

```go
import (
	"fmt"
	"uniqid" // your local uniqid package from your project
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(uniqid.Generate())
	}
}
```

### Using with Dependency Injection

```go
package main

// Import packages
import (
	"fmt"
	"github.com/mustafaturan/monoton/v3"
	"github.com/mustafaturan/monoton/v3/sequencer"
)

func NewIDGenerator() monoton.Monoton {
	// Fetch your node id from a config server or generate from MAC/IP address
	node := uint64(1)

	// A unix time value which will be subtracted from the time sequence value.
	// The initialTime value type corresponds to the sequencer type's time
	// representation. If you are using Millisecond sequencer then it must be
	// considered as Millisecond
	initialTime := uint64(0)

	// Configure monoton with a sequencer and the node
	m, err := monoton.New(sequencer.NewMillisecond(), node, initialTime)
	if err != nil{
		panic(err)
	}

	return m
}

func main() {
	g := NewIDGenerator()

	for i := 0; i < 100; i++ {
		fmt.Println(g.Next())
	}
}
```

## Features

### Time Ordered

The `monoton` package provides sequences based on the `monotonic` time which
represents the absolute elapsed wall-clock time since some arbitrary, fixed
point in the past. It isn't affected by changes in the system time-of-day clock.

Please refer to [ADR 01 - Time](docs/adrs/time.md) for details and consequences.

### Initial Time

Initial time value opens space for time value by subtracting the given value
from the time sequence.

### Readable

The `monoton` package converts all sequences into `Base62` format. And `Base62`
only uses `ASCII` alpha-numeric chars to represent data which makes it easy to
read, predict the order by a human eye.

The total byte size is fixed to 16 bytes for all sequencers. And at least one
byte is reserved to nodes.

Please refer to [ADR 02 - Encoding](docs/adrs/encoding.md) for details and
consequences.

### Multi Node Support

The `monoton` package can be used on single/multiple nodes without the need for
machine coordination. It uses configured node identifier to generate ids by
attaching the node identifier to the end of the sequences.

### Extendable

The package comes with three pre-configured sequencers and `Sequencer` interface
to allow new sequencers.

#### Included Sequencers and Byte Orderings

The `monoton` package currently comes with `Nanosecond`, `Millisecond` and
`Second` sequencers. And it uses `Millisecond` sequencer by default. For each
sequencer, the byte orders are as following:

```
Second:      16 B =>  6 B (seconds)      + 6 B (counter) + 4 B (node)
Millisecond: 16 B =>  8 B (milliseconds) + 4 B (counter) + 4 B (node)
Nanosecond:  16 B => 11 B (nanoseconds)  + 2 B (counter) + 3 B (node)
```

Please refer to [ADR 03 - Byte Sizes](docs/adrs/byte-sizes.md) for details and
consequences.

#### New Sequencers

The sequencers can be extended for any other time format, sequence format by
implementing the `monoton/sequencer.Sequencer` interface.

## Benchmarks

Command:
```
go test -benchtime 10000000x -benchmem -run=^$ -bench=. github.com/mustafaturan/monoton/v3
```

Results:
```
goos: darwin
goarch: amd64
pkg: github.com/mustafaturan/monoton/v3
cpu: Intel(R) Core(TM) i5-6267U CPU @ 2.90GHz
BenchmarkNext-4        	10000000	       102.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkNextBytes-4   	10000000	        97.51 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/mustafaturan/monoton/v3	2.203s
```

## Contributing

All contributors should follow [Contributing Guidelines](CONTRIBUTING.md) and
[ADR docs](docs/adrs) before creating pull requests.

## Credits

[Mustafa Turan](https://github.com/mustafaturan)

## License

Apache License 2.0

Copyright (c) 2019 Mustafa Turan

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
