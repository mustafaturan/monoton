package monoton_test

import (
	"fmt"
	"time"

	"github.com/mustafaturan/monoton"
	"github.com/mustafaturan/monoton/sequencer"
)

func ExampleConfigure() {
	s := sequencer.NewMillisecond() // has 4 bytes free space for a node
	n := uint64(19)                 // Base62 => J
	t := uint64(0)                  // initial time (start from unix time in ms)

	monoton.Configure(s, n, t)
	fmt.Println(monoton.Next()[12:])
	// Output:
	// 000J
}

func ExampleNext() {
	s := sequencer.NewSecond()     // sequencer.Second
	n := uint64(19)                // Base62 => J
	t := uint64(time.Now().Unix()) // initial time (start from unix time in s)

	monoton.Configure(s, n, t)
	fmt.Println(len(monoton.Next()))
	// Output:
	// 16
}
