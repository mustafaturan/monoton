# Monoton

[![Build Status](https://travis-ci.org/mustafaturan/monoton.svg?branch=master)](https://travis-ci.org/mustafaturan/monoton)

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
	monoton.Configure(sequencer.NewMillisecond(), node)
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(monoton.Next())
	}
}
```

### New Sequencers

The `monoton` package currently comes with `Nanosecond`, `Millisecond` and
`Second` sequencers. And it uses `Millisecond` sequencer by default.
But, for sure depending on needs, it can be extended for `Microsecond` easily by
implementing the `monoton.Sequencer` interface.

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
