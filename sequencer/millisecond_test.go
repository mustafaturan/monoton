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
