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

func TestMultiple32(t *testing.T) {
	var s = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdef")
	if len(s) != 32 {
		t.Fatal("s needs to be 32")
	}
	x := rand_char(32, s)
	if len(x) != 32 {
		t.Fatal("The password returned was too short")
	}
}
