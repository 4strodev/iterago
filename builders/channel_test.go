package builders_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/4strodev/iterago/builders"
	"github.com/4strodev/iterago/transformers"
	"github.com/stretchr/testify/assert"
)

func TestChannel(t *testing.T) {
	dataChannel := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			dataChannel <- i
		}
		close(dataChannel)
	}()

	channelIterator := builders.Channel(dataChannel)
	for i, v := range channelIterator {
		assert.True(t, i >= 0 && i < 10, "index should be between 0 and 10 i=%d", i)
		assert.True(t, v >= 0 && v < 10, "value should be between 0 and 10 v=%d", v)
	}
}

func TestChannelPanic(t *testing.T) {
	dataChannel := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			dataChannel <- i
		}
		close(dataChannel)
	}()

	channelIterator := builders.Channel(dataChannel)
	for _, v := range channelIterator {
		if v > 5 {
			break
		}
	}
}

func ExampleChannel() {
	channel := make(chan int)
	go func() {
		for _, v := range builders.Range(0, 10, 1) {
			channel <- v
		}
		close(channel)
	}()

	iterator := builders.Channel(channel)
	evenIterator := transformers.Filter(iterator, func(i int, v int) bool {
		return v%2 == 0
	})
	evenValues := slices.Collect(transformers.Values(evenIterator))
	fmt.Println(evenValues)
	// Output: [0 2 4 6 8]
}
