package builders

import "iter"

// Channel allows you to convert a channel into an interator
func Channel[V any](channel chan V) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		index := 0
		for v := range channel {
			if !yield(index, v) {
				return
			}
			index++
		}
	}
}
