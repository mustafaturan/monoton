// Copyright 2021 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package monoton_test

import (
	"fmt"
	"time"

	"github.com/mustafaturan/monoton/v2"
	"github.com/mustafaturan/monoton/v2/sequencer"
)

func ExampleNew() {
	s := sequencer.NewMillisecond() // has 4 bytes free space for a node
	n := uint64(19)                 // Base62 => J
	t := uint64(0)                  // initial time (start from unix time in ms)

	m, err := monoton.New(s, n, t)
	if err != nil {
		panic(err)
	}
	fmt.Println(m.Next()[12:])
	// Output:
	// 000J
}

func ExampleMonoton_Next() {
	s := sequencer.NewSecond()     // sequencer.Second
	n := uint64(19)                // Base62 => J
	t := uint64(time.Now().Unix()) // initial time (start from unix time in s)

	m, err := monoton.New(s, n, t)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(m.Next()))
	// Output:
	// 16
}
