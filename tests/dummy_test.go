package tests

import "testing"

//fill this code file with some dummy tests

func TestDummyTestReturnsTrue(t *testing.T) {
	if true != false {
		t.Error("DummyTestReturnsTrue")
	} else {
		t.Log("Success")
	}
}
