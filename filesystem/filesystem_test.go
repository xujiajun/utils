package filesystem

import "testing"

func TestPathIsExist(t *testing.T) {
	expected := false
	if PathIsExist("/Users/xujiajun123") != expected {
		t.Errorf("returned unexpected bool value : got %v want %v", !expected, expected)
	}
}
