package builders_test

import (
	"slices"
	"testing"

	"github.com/4strodev/iterators/builders"
	"github.com/4strodev/iterators/transformers"
)

func bencharkRange(start int, end int, jump int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		iterator := builders.Range(start, end, jump)
		slices.Collect(transformers.Values(iterator))
	}
}

func BenchmarkRange10(b *testing.B) {
	bencharkRange(0, 10, 1, b)
}

func BenchmarkRange100(b *testing.B) {
	bencharkRange(0, 100, 1, b)
}

func BenchmarkRange1000(b *testing.B) {
	bencharkRange(0, 1000, 1, b)
}
