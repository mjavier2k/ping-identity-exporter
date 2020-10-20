package prom

import (
	"testing"
)

func TestConverToBytes(t *testing.T) {

	var tests = []struct {
		val      float64
		uom      string
		expected float64
	}{
		{1, "MB", 1000000},
		{1, "GB", 1000000000},
		{1, "", 1},
	}

	for _, test := range tests {
		if output := ConvertToBytes(test.val, test.uom); output != test.expected {
			t.Error("Test Failed. {} input, {} expected, received {}", test.val, test.expected, output)
		}
	}
}
