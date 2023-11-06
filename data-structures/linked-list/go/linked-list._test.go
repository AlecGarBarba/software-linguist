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

	for value := range values {
		sll.Append(value)
	}

	assert.Equal(t, sll.Length(), uint(len(values)))

}
