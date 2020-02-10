// Copyright 2020 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"reflect"
	"testing"
)

func TestNewNanosecond(t *testing.T) {
	want := reflect.TypeOf(&Sequence{})
	got := reflect.TypeOf(NewNanosecond())

	if want != got {
		t.Errorf("NewNanosecond() call want type: %T, got type: %T", want, got)
	}
}
