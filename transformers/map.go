package transformers

import "iter"

// Map allows you to convert an iterator [iter.Seq2] into another iterator with mapped keys and values
func Map[IK any, IV any, OK any, OV any](iterator iter.Seq2[IK, IV], mapper func(k IK, v IV) (OK, OV)) iter.Seq2[OK, OV] {
	return func(yield func(OK, OV) bool) {
		for k, v := range iterator {
			key, value := mapper(k, v)
			if !yield(key, value) {
				return
			}
		}
	}
}
