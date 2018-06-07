package stringutil

import (
	"testing"
)

func TestReverse(t *testing.T) {
	test_cases := []struct {
		in, want string
	}{
		{"Go's testing package.", ".egakcap gnitset s'oG"},
		{"This is 世界", "界世 si sihT"},
		{"Ewan Ewart", "trawE nawE"},
		{"stringutil", "litugnirts"},
	}
	for _, c := range test_cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
