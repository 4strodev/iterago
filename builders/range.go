package builders

import "iter"

// Range returns an iterator with the specified boundaries
// it starts the iterator with an iteration value that is equal to start and ends when the iteration value is equal to end.
// After each iteration the next value is calculated summing jump to the current value.
// If the parameters are put correctly an infinite or reverse iterator can be created.
func Range(start int, end int, jump int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		var value = start
		var index = 0
		for {
			if value == end {
				return
			}
			if !yield(index, value) {
				return
			}
			index += 1
			value += jump
		}
	}
}
