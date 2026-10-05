package randstr

import "testing"

func TestStringEmptyAlphabet(t *testing.T) {
	if got := String(8, ""); got != "" {
		t.Fatalf("String(8, \"\") = %q, want empty", got)
	}
}

func TestStringNonPositive(t *testing.T) {
	if got := String(0); got != "" {
		t.Fatalf("String(0) = %q, want empty", got)
	}
	if got := String(-3); got != "" {
		t.Fatalf("String(-3) = %q, want empty", got)
	}
}

func TestBytesNonPositive(t *testing.T) {
	if got := Bytes(0); len(got) != 0 {
		t.Fatalf("Bytes(0) len = %d, want 0", len(got))
	}
	if got := Bytes(-1); len(got) != 0 {
		t.Fatalf("Bytes(-1) len = %d, want 0", len(got))
	}
}

func TestStringStillWorks(t *testing.T) {
	s := String(16, "abc")
	if len(s) != 16 {
		t.Fatalf("len = %d, want 16", len(s))
	}
	for _, r := range s {
		if r != 'a' && r != 'b' && r != 'c' {
			t.Fatalf("unexpected rune %q", r)
		}
	}
}
