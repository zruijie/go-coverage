package mymath

import "testing"

func TestIsBigNum(t *testing.T) {
	if !IsBigNum(1001) {
		t.Error("1001 should be a big number")
	}

	if IsBigNum(999) {
		t.Error("999 should not be a big number")
	}
}

func TestIsSmallNum(t *testing.T) {
	// if !IsSmallNum(999) {
	// 	t.Error("999 should be a small number")
	// }

	if IsSmallNum(1001) {
		t.Error("1001 should not be a small number")
	}
}
