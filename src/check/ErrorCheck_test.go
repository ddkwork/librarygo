package check

import "testing"

func TestCheck(t *testing.T) {
	check := New()
	check.Error("Π√√√√√√√√√√√√√Π√√√√√√√√√√√√√Π√√√√√√√√√√√√√Π√√√√√√√√√√√√√Π√√√√√√√√√√√√√")
}
