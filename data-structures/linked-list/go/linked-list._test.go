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

func TestSinglyLinkedListInsertAt(t *testing.T) {

	values := []int{1, 2, 3, 4, 5}
	sll := SingleLinkedList[int]{}

	for value := range values {
		sll.Append(value)
	}
	/*insert out of index range*/

	err := sll.InsertAt(6, len(values)+5)

	assert.Error(t, err, "Error expected as index to insert is out of range")
	/*insert beginning*/
	sll.InsertAt(0, 0)
	head, _ := sll.Get(0)
	assert.Equal(t, *head, 0)

	/*insert at the end*/

	end_err := sll.InsertAt(6, 6)
	if end_err != nil {
		t.Errorf("Error inserting at end of linked list, reason: %v", end_err)
	}
	tail, _ := sll.Get(6)

	assert.Equal(t, *tail, 6)
}

func TestSinglyLinkedListInsertAtHead(t *testing.T) {
	sll := SingleLinkedList[int]{}

	index_out_of_range_err := sll.InsertAt(10, 10)

	assert.Error(t, index_out_of_range_err, "Should error on insertAt without head")

	head_err := sll.InsertAt(0, 0)

	assert.Nil(t, head_err, "Should be able to insert head")
}

func TestSinglyLinkedList(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}

	sll := SingleLinkedList[int]{}
	for _, value := range values {
		sll.Append(value)
	}

	/*remove at the end*/
	end_val, end_err := sll.RemoveAt(4)
	assert.Nil(t, end_err, "Should properly remove last index")
	assert.Equal(t, 5, *end_val)
	/*remove at beggining*/
	beg_val, beg_err := sll.RemoveAt(0)
	/*Remove until list is empty*/
	assert.Nil(t, beg_err, "Should properly remove at beggining")
	assert.Equal(t, values[0], *beg_val)

	for i := 0; i < int(sll.Length()); i++ {
		val, err := sll.RemoveAt(0)

		assert.Nil(t, err, "Remove out of range")

		assert.NotNil(t, val)

	}
}
