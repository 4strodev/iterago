package builders_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/4strodev/iterago/builders"
	"github.com/4strodev/iterago/transformers"
	"github.com/stretchr/testify/assert"
)

func TestRangeBuilderBreak(t *testing.T) {
	start := 0
	end := -1
	jump := 1

	for i, _ := range builders.Range(start, end, jump) {
		if i == 100 {
			break
		}
	}
}

func TestRangeBuilderConsistency(t *testing.T) {
	start := 0
	end := -1
	jump := 1

	iterator := builders.Range(start, end, jump)

	limit := 1000
	indexCounter := 0
	valueCounter := start

	for i, v := range iterator {
		if indexCounter == limit {
			break
		}

		if indexCounter == 0 {
			assert.Equal(t, start, v, "it should start at %d", start)
		}

		a := assert.Equal(t, indexCounter, i, "they should be equal counter=%d params=%v", indexCounter, []int{start, end, jump})
		b := assert.Equal(t, valueCounter, v, "they should be equal counter=%d params=%v", indexCounter, []int{start, end, jump})
		if !a || !b {
			break
		}

		indexCounter++
		valueCounter += jump
	}
}

func ExampleRange() {
	rangeIterator := builders.Range(0, 10, 2)

	slice := slices.Collect(transformers.Values(rangeIterator))
	fmt.Println(slice)
	// Output: [0 2 4 6 8]

}

func ExampleRange_infinite() {
	// Infinite loop
	for _, v := range builders.Range(0, -1, 1) {
		if v == 10 {
			break
		}
		fmt.Println(v)
	}
	/* Output:
0
1
2
3
4
5
6
7
8
9
	*/
}

func ExampleRange_reverse() {
	rangeIterator := builders.Range(10, 0, -2)

	slice := slices.Collect(transformers.Values(rangeIterator))
	fmt.Println(slice)
	// Output: [10 8 6 4 2]
}
