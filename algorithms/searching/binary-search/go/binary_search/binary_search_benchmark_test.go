package binarysearch

import (
	"testing"
)

func createOrderedArray(size uint32) []uint32 {
	a := make([]uint32, size)
	for i := range a {
		a[i] = uint32(i)
	}
	return a
}

var array_10k = createOrderedArray(10_000)
var array_100k = createOrderedArray(100_000)
var array_1M = createOrderedArray(1_000_000)

func BenchmarkBinarySearch10K(b *testing.B) {
	for i := 0; i < 10_000; i++ {

		Binarysearch(&array_10k, uint32(i))
	}
}

func BenchmarkBinarySearch100K(b *testing.B) {
	for i := 0; i < 100_000; i++ {
		Binarysearch(&array_100k, uint32(i))
	}
}

func BenchmarkBinarySearch1M(b *testing.B) {
	for i := 0; i < 1_000_000; i++ {
		Binarysearch(&array_1M, uint32(i))
	}
}
