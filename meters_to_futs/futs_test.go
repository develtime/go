package main

import "testing"

func TestMeterToFuts(t *testing.T) {
	futs := MeterToFuts(3048)

	if futs != 10000 {
		t.Error("Expected 10000, got ", futs)
	}
}
