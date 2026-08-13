package gonanoid

import (
	"math"
	"testing"
)

func TestGetMask(t *testing.T) {
	tests := []struct {
		alphabetSize int
		want         int
	}{
		{alphabetSize: 1, want: 1},
		{alphabetSize: 2, want: 1},
		{alphabetSize: 3, want: 3},
		{alphabetSize: 64, want: 63},
		{alphabetSize: 65, want: 127},
		{alphabetSize: 256, want: 255},
	}

	for _, tt := range tests {
		if got := getMask(tt.alphabetSize); got != tt.want {
			t.Fatalf("getMask(%d) = %d, want %d", tt.alphabetSize, got, tt.want)
		}
	}
}

func TestHasNoCollisions(t *testing.T) {
	tries := 100_000
	used := make(map[string]bool)
	for i := 0; i < tries; i++ {
		id := Must()
		if used[id] {
			t.Fatalf("generated colliding ID %q", id)
		}
		used[id] = true
	}
}

func TestFlatDistribution(t *testing.T) {
	tries := 100_000
	alphabet := "abcdefghij"
	size := 10
	chars := make(map[rune]int)
	for i := 0; i < tries; i++ {
		id := MustGenerate(alphabet, size)
		for _, r := range id {
			chars[r]++
		}
	}

	for _, count := range chars {
		want := float64(size * tries / len(alphabet))
		if delta := math.Abs(float64(count)-want) / want; delta > 0.01 {
			t.Fatalf("character count %d differs from expected %.0f by %.2f%%", count, want, delta*100)
		}
	}
}

// Benchmark nanoid generator
func BenchmarkNanoid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = New()
	}
}
