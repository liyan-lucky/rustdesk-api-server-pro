package util

import "testing"

func TestConstantTimeStringEqual(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{name: "equal", a: "same-secret", b: "same-secret", want: true},
		{name: "different same length", a: "same-secret", b: "same-secreu", want: false},
		{name: "different length", a: "short", b: "shorter", want: false},
		{name: "empty equal", a: "", b: "", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConstantTimeStringEqual(tt.a, tt.b); got != tt.want {
				t.Fatalf("ConstantTimeStringEqual(%q, %q) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
