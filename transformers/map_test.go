package transformers_test

import (
	"fmt"
	"slices"
	"strconv"
	"testing"

	"github.com/4strodev/iterators/builders"
	"github.com/4strodev/iterators/transformers"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	iterator := transformers.Map(builders.Range(0, 10, 1), func(index int, value int) (string, int) {
		return strconv.Itoa(index), value * 2
	})

	values := transformers.Values(iterator)
	for v := range values {
		if !assert.Equal(t, v%2 == 0, true, "value should be pair v=%d", v) {
			break
		}
	}
}

func TestMapBreak(t *testing.T) {
	iterator := transformers.Map(builders.Range(0, 10, 1), func(index int, value int) (int, int) {
		return index, value
	})

	values := transformers.Values(iterator)
	for v := range values {
		if v > 5 {
			break
		}
	}
}

func ExampleMap() {
	iterator := transformers.Map(builders.Range(0, 10, 1), func(index int, value int) (string, int) {
		return fmt.Sprintf("\"%d\"", index), value * 2
	})

	values := transformers.Values(iterator)
	valueItems := slices.Collect(values)

	keys := transformers.Keys(iterator)
	keyItems := slices.Collect(keys)

	fmt.Println(valueItems)
	fmt.Println(keyItems)
	// Output:
	// [0 2 4 6 8 10 12 14 16 18]
	// ["0" "1" "2" "3" "4" "5" "6" "7" "8" "9"]
}
