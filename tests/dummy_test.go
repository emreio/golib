package tests

import "testing"

func TestDummyTestReturnsTrue(t *testing.T) {
	if true != false {
		t.Error("DummyTestReturnsTrue")
	} else {
		t.Log("Success")
	}
}
