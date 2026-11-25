package vo

import "testing"

func TestNumberValid(t *testing.T) {
	n, err := NewNumber("12.5")
	if err != nil || n.String() != "12.5" {
		t.Fatal(err)
	}
}

func TestNumberInvalid(t *testing.T) {
	if _, err := NewNumber("0"); err == nil {
		t.Fatal("expected error")
	}
	if _, err := NewNumber("abc"); err == nil {
		t.Fatal("expected error")
	}
}
