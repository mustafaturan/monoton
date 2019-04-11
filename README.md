# Monoton

[![Build Status](https://travis-ci.org/mustafaturan/monoton.svg?branch=master)](https://travis-ci.org/mustafaturan/monoton)
[![Coverage Status](https://coveralls.io/repos/github/mustafaturan/monoton/badge.svg?branch=master)](https://coveralls.io/github/mustafaturan/monoton?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mustafaturan/monoton)](https://goreportcard.com/report/github.com/mustafaturan/monoton)
[![GoDoc](https://godoc.org/github.com/mustafaturan/monoton?status.svg)](https://godoc.org/github.com/mustafaturan/monoton)

Highly scalable, single/multi node, predictable and incremental unique id
generator.

## Installation

Via go packages:
```go get github.com/mustafaturan/monoton```

## Usage

```go
package main

// Import packages
import (
	"fmt"
	"github.com/mustafaturan/monoton"
	"github.com/mustafaturan/monoton/sequencer"
)

// On init configure the monoton
func init() {
	// Fetch your node id from a config server or generate from MAC/IP address
	node := uint(1)

	// A unix time value which will be subtracted from the time sequence value.
	// The initialTime value type corresponds to the sequencer type's time
	// representation. If you are using Millisecond sequencer then it must be
	// considered as Millisecond
	initialTime := uint(0)

	// Configure monoton with a sequencer and the node
	monoton.Configure(sequencer.NewMillisecond(), node, initialTime)
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(monoton.Next())
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
