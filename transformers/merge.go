package transformers

import (
	"context"
	"iter"
)

// Merge allows you to merge multiple iterators into one.
func Merge[K any, V any](iterators ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	type Pairs struct {
		key   K
		value V
	}

	return func(yield func(K, V) bool) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		endSignal := make(chan struct{})

		entriesChannel := make(chan Pairs)

		// this goroutine handles iterators finishes
		// when both iterators are finished then it closes the entries channel
		go func() {
			counter := 0
			for range endSignal {
				counter++
				if counter == len(iterators) {
					close(entriesChannel)
				}
			}
		}()

		for _, iterator := range iterators {
			go func() {
				defer func() {
					endSignal <- struct{}{}
				}()
				for k, v := range iterator {
					select {
					case <-ctx.Done():
						return
					default:
						entriesChannel <- Pairs{k, v}
					}
				}
			}()
		}

		for entry := range entriesChannel {
			if !yield(entry.key, entry.value) {
				return
			}
		}
	}

}
