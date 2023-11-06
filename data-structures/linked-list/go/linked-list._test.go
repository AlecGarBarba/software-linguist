package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
TESTING SINGLY LINKED LIST
*/

func TestSinglyLInkedListAppend(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	sll := SingleLinkedList[int]{}
	for index, value := range values {
		sll.Append(value)
		ll_value, err := sll.Get(index)
		if err != nil {
			t.Errorf("Error appending and getting last value")
		}
		assert.Equal(t, value, *ll_value)
	}
	assert.Equal(t, sll.Length(), uint(len(values)))

}

func TestSinglyLInkedListPrepend(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	sll := SingleLinkedList[int]{}
	for value := range values {
		sll.Prepend(value)
		head, err := sll.Get(0)
		if err != nil {
			t.Errorf("Error prepending and getting first ealement")
		}
		assert.Equal(t, *head, value)
	}
	assert.Equal(t, sll.Length(), uint(len(values)))
}
