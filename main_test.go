package main

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestStringsRepeat(t *testing.T) {
	got := StringsRepeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkStringsRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringsRepeat("a", 5)
	}
}
