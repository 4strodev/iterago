package transformers

import "iter"

// Filter allows to filter values of an interator specifying a predicate that gets
func Filter[K any, V any](iterator iter.Seq2[K, V], predicate func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range iterator {
			if !predicate(k, v) {
				continue
			}

			if !yield(k, v) {
				return
			}
		}
	}
}
