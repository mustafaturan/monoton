// Copyright 2020 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"reflect"
	"testing"
)

func TestNewSecond(t *testing.T) {
	want := reflect.TypeOf(&Sequence{})
	got := reflect.TypeOf(NewSecond())

	if want != got {
		t.Errorf("NewSecond() call want type: %T, got type: %T", want, got)
	}
}
