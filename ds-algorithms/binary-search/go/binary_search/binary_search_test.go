package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchEmptyArray(t *testing.T) {
	emptyArray := []uint32{}
	assert.False(t, Binarysearch(emptyArray, 0), "True is true!")

}

func TestBinarySearchValueNotInArray(t *testing.T) {
	even_array := []uint32{2, 4, 6, 8, 10, 12, 14, 16}

	assert.False(t, Binarysearch(even_array, 1), "Should return false if value does not exist in array")
}

func TestBinarySearchValueStartArray(t *testing.T) {
	simple_array := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.True(t, Binarysearch(simple_array, simple_array[0]), "Should return true if element is at the start of the array")
}

func TestBinarySearchValueEndArray(t *testing.T) {
	simple_array := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.True(t, Binarysearch(simple_array, simple_array[9]), "Should return true if element is at the start of the array")
}
func TestBinarySearchValueMiddleArray(t *testing.T) {
	simple_array := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.True(t, Binarysearch(simple_array, simple_array[5]), "Should return true if element is at the start of the array")
}

//Colorful outputs https://stackoverflow.com/questions/27242652/colorizing-golang-test-run-output
