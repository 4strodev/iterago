package transformers_test

import (
	"testing"

	"github.com/4strodev/iterators/builders"
	"github.com/4strodev/iterators/transformers"
	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	values := transformers.Values(builders.Range(0, 10, 1))

	counter := 0
	for v := range values {
		assert.Equal(t, counter, v, "value should be equal to counter value=%d counter=%d", v, counter)
		counter ++
	}
}

func TestValueBreak(t *testing.T) {
	values := transformers.Values(builders.Range(0, 10, 1))

	for v := range values {
		if v > 5 {
			break
		}
	}
}

func TestKeys(t *testing.T) {
	keys := transformers.Keys(builders.Range(0, 10, 1))

	counter := 0
	for k := range keys {
		assert.Equal(t, counter, k, "key should be equal to counter key=%d counter=%d", k, counter)
		counter ++
	}
}

func TestKeysBreak(t *testing.T) {
	keys := transformers.Keys(builders.Range(0, 10, 1))

	for k := range keys {
		if k > 5 {
			break
		}
	}
}
