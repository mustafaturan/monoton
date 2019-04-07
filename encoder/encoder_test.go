package encoder

import "testing"

func TestToBase62(t *testing.T) {
	tests := []struct {
		val         uint64
		expectedVal string
	}{
		{1, "1"},
		{63, "11"},
		{1<<64 - 1, "LygHa16AHYF"},
	}

	msg := "ToBase62(%d) = %v, but returned %v"
	for _, test := range tests {
		result := ToBase62(test.val)
		if result != test.expectedVal {
			t.Errorf(msg, test.val, test.expectedVal, result)
		}
	}
}

func TestToBase62WithPaddingZeros(t *testing.T) {
	tests := []struct {
		val         uint
		padding     int64
		expectedVal string
	}{
		{1, 11, "00000000001"},
		{63, 11, "00000000011"},
		{1<<64 - 1, 11, "LygHa16AHYF"},
	}

	msg := "ToBase62WithPaddingZeros(%d, %d) = %v, but returned %v"
	for _, test := range tests {
		result := ToBase62WithPaddingZeros(test.val, test.padding)
		if result != test.expectedVal {
			t.Errorf(msg, test.val, test.padding, test.expectedVal, result)
		}
	}
}
