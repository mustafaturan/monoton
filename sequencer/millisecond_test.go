// Copyright 2021 Mustafa Turan. All rights reserved.
// Use of this source code is governed by a Apache License 2.0 license that can
// be found in the LICENSE file.

package sequencer

import (
	"reflect"
	"testing"
)

func TestNewMillisecond(t *testing.T) {
	want := reflect.TypeOf(&Sequence{})
	got := reflect.TypeOf(NewMillisecond())

	if want != got {
		t.Errorf("NewMillisecond() call want type: %T, got type: %T", want, got)
	}
}
