package math

import "testing"

func TestMeterToFuts(t *testing.T) {
	value = MeterToFuts(3048)

	if value != 10000 {
		t.Error("Expected 10000, got ", value)
	}
}
