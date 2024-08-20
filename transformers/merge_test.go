package transformers_test

import (
	"slices"
	"testing"

	"github.com/4strodev/iterago/builders"
	"github.com/4strodev/iterago/transformers"
	"github.com/stretchr/testify/assert"
)

func TestMergeIterators(t *testing.T) {
	iter1 := builders.Range(0, 5, 1)
	iter2 := builders.Range(5, 10, 1)

	iter3 := transformers.Merge(iter1, iter2)
	values := slices.Collect(transformers.Values(iter3))
	assert.Equal(t, 10, len(values))
}

func TestMergeBreak(t *testing.T) {
	iter1 := builders.Range(0, 5, 1)
	iter2 := builders.Range(5, 10, 1)

	iter3 := transformers.Merge(iter1, iter2)
	for _, v := range iter3 {
		if v > 5 {
			break
		}
	}
}
