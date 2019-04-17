package mtimer

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	t.Run("in a future time", func(t *testing.T) {
		t1 := Now()
		time.Sleep(time.Nanosecond)
		t2 := Now()
		if t1 >= t2 {
			t.Errorf(
				"Now() after enough sleep should increment %d = %d",
				t1,
				t2,
			)
		}
	})
}
