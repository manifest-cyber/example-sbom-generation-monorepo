package dummy

import (
	"testing"
)

func TestDummy(t *testing.T) {
	result := Dummy("works")
	if result != "Dummy works" {
		t.Error("Expected Dummy to append 'works'")
	}
}
