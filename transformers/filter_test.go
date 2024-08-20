package transformers_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/4strodev/iterago/builders"
	"github.com/4strodev/iterago/transformers"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	iterator := transformers.Filter(builders.Range(0, 10, 1), func(k int, v int) bool {
		return v%2 == 0
	})

	items := transformers.Values(iterator)
	for v := range items {
		if !assert.Equal(t, v%2 == 0, true, "should be pair v=%d", v) {
			return
		}
	}
}

func TestFilterBreak(t *testing.T) {
	iterator := transformers.Filter(builders.Range(0, 10, 1), func(k int, v int) bool {
		return true
	})

	items := transformers.Values(iterator)
	for v := range items {
		if v > 5 {
			break
		}
	}
}

func ExampleFilter() {
	iterator := transformers.Filter(builders.Range(0, 10, 1), func(k int, v int) bool {
		return v%2 == 0
	})

	items := slices.Collect(transformers.Values(iterator))
	fmt.Println(items)
	//Output: [0 2 4 6 8]
}
