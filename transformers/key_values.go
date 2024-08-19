package transformers

import "iter"

// Keys returns the keys of a iterator [iter.Seq2]
func Keys[K any, V any](iterator iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range iterator {
			if !yield(k) {
				return
			}
		}
	}
}

// Values returns the values of a iterator [iter.Seq2]
func Values[K any, V any](iterator iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range iterator {
			if !yield(v) {
				return
			}
		}
	}
}
