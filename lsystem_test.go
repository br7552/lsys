package lsystem

import "testing"

func TestLsystem(t *testing.T) {
	l := NewLsystem("B", "B", "F[-B]+B", "F", "FF")

	assertString(t, "B", l.String())

	for i := 0; i < 3; i++ {
		l.Grow()
	}

	want := "FFFF[-FF[-F[-B]+B]+F[-B]+B]+FF[-F[-B]+B]+F[-B]+B"
	assertString(t, want, l.String())
}

var blackhole string

func BenchmarkLSystem(b *testing.B) {
	l := NewLsystem("F", "F", "FF+[+F-F-F]-[-F+F+F]")

	for i := 0; i < 10; i++ {
		l.Grow()
	}

	blackhole = l.String()
}

func assertString(t *testing.T, want, got string) {
	t.Helper()

	if got != want {
		t.Fatalf("expected: %q, got %q", want, got)
	}
}
