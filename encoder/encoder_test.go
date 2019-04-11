package encoder

import "testing"

func TestToBase62(t *testing.T) {
	tests := []struct {
		val  uint64
		want string
	}{
		{1, "1"},
		{63, "11"},
		{1<<64 - 1, "LygHa16AHYF"},
	}

	msg := "ToBase62(%d) = %v, but returned %v"
	for _, test := range tests {
		got := ToBase62(test.val)
		if got != test.want {
			t.Errorf(msg, test.val, test.want, got)
		}
	}
}

func TestToBase62WithPaddingZeros(t *testing.T) {
	tests := []struct {
		val     uint
		padding int64
		want    string
	}{
		{1, 11, "00000000001"},
		{63, 11, "00000000011"},
		{1<<64 - 1, 11, "LygHa16AHYF"},
	}

	msg := "ToBase62WithPaddingZeros(%d, %d) = %v, but returned %v"
	for _, test := range tests {
		got := ToBase62WithPaddingZeros(test.val, test.padding)
		if got != test.want {
			t.Errorf(msg, test.val, test.padding, test.want, got)
		}
	}
}
