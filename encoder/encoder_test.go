package encoder

import "testing"

func TestToBase62WithPaddingZeros(t *testing.T) {
	tests := []struct {
		val     uint64
		padding int
		want    string
	}{
		{1, 11, "00000000001"},
		{63, 2, "11"},
		{124, 3, "020"},
		{125, 4, "0021"},
		{1<<64 - 1, 11, "LygHa16AHYF"},
	}

	msg := "ToBase62WithPaddingZeros(%d, %d) = %v, but returned %v"
	for _, test := range tests {
		got := ToBase62WithPaddingZeros(test.val, test.padding)
		if string(got) != test.want {
			t.Errorf(msg, test.val, test.padding, test.want, string(got))
		}
	}
}

func TestBase62ByteSize(t *testing.T) {
	tests := []struct {
		val  uint64
		want int
	}{
		{0, 1},
		{1, 1},
		{63, 2},
		{124, 2},
		{125, 2},
		{1<<64 - 1, 11},
	}

	msg := "Base62ByteSize(%d) = %d, but returned %v"
	for _, test := range tests {
		got := Base62ByteSize(test.val)
		if got != test.want {
			t.Errorf(msg, test.val, test.want, got)
		}
	}
}
