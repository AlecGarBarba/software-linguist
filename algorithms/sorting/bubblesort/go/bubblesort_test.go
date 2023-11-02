package bubblesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortAscending(t *testing.T) {

	arr := []int{1, 3, 7, 2, 4, 6, 5, 8}

	BubbleSort(&arr, Ascending)

	assert.Equal(t, &arr, &[]int{1, 2, 3, 4, 5, 6, 7, 8}, "Array is not ordererd in assencing order")

}

func TestBubbleSortDescending(t *testing.T) {

	arr := []int{1, 3, 7, 2, 4, 6, 5, 8}

	BubbleSort(&arr, Descending)

	assert.Equal(t, &arr, &[]int{8, 7, 6, 5, 4, 3, 2, 1}, "Array is not ordererd in assencing order")

}

func TestBubbleSortEmpty(t *testing.T) {
	empty_arr := []int{}

	assert.Equal(t, &empty_arr, &[]int{})
}
