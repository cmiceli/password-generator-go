package pwordgen

import (
	"testing"
)

func TestNewPassword(t *testing.T) {
	for i := 0; i < len(StdChars); i++ {
		x := NewPassword(i)
		if len(x) != i {
			t.Fatal("The password returned was too short")
		}
	}
}

func TestNewPasswordMoreThanStdChars(t *testing.T) {
	for i := len(StdChars); i < len(StdChars)*2; i++ {
		x := NewPassword(i)
		if len(x) != i {
			t.Fatal("The password returned was too short")
		}
	}
}
