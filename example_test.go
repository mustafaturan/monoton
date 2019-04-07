package monoton_test

import (
	"fmt"

	"github.com/mustafaturan/monoton"
	"github.com/mustafaturan/monoton/sequencer"
)

func ExampleConfigure() {
	s := sequencer.NewMillisecond() // has 4 bytes free space for a node
	n := uint(19)                   // Base62 => J

	monoton.Configure(s, n)
	fmt.Println(monoton.Next()[12:])
	// Output:
	// 000J
}

func ExampleNext() {
	s := sequencer.NewMillisecond()
	n := uint(19)

	monoton.Configure(s, n)
	fmt.Println(len(monoton.Next()))
	// Output:
	// 16
}
