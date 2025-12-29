package mathutils

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestSub(t *testing.T) {
	result := Sub(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}
