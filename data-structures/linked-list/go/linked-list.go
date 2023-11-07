package linked_list

import (
	"fmt"
)

type ILinkedList[T interface{}] interface {
	Length() uint
	InsertAt(item T, index int) error
	RemoveAt(index int) (T, error)
	Append(item T)
	Prepend(item T)
	Get(index int) (T, error)
}

type Node[T interface{}] struct {
	value T
	next  *Node[T]
}

type SingleLinkedList[T interface{}] struct {
	head *Node[T]
	len  uint
}

func (ll *SingleLinkedList[T]) Length() uint {
	return ll.len
}

func (ll *SingleLinkedList[T]) Append(item T) {
	newNode := &Node[T]{value: item}
	ll.len++
	if ll.head == nil {
		ll.head = newNode
		return
	}
	curr := ll.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (ll *SingleLinkedList[T]) Prepend(item T) {
	newNode := &Node[T]{value: item}
	ll.len++

	if ll.head == nil {
		ll.head = newNode
		return
	}

	newNode.next = ll.head
	ll.head = newNode
}

func (ll *SingleLinkedList[T]) Get(index int) (*T, error) {
	curr := ll.head

	for i := 0; i < index; i++ {
		if curr == nil {
			return nil, fmt.Errorf("Value not found")
		}
		curr = curr.next
	}

	return &curr.value, nil
}

func (ll *SingleLinkedList[T]) InsertAt(item T, index int) error {
	if ll.Length() == 0 && index == 0 {
		ll.head = &Node[T]{value: item}
		ll.len++
		return nil
	}
	curr := ll.head
	for i := 1; i < index; i++ {
		if curr == nil {
			return fmt.Errorf("Cannot insert at index %v since it's larged than current length %v", index, ll.len)
		}
		curr = curr.next
	}
	new := &Node[T]{value: item, next: curr.next}
	curr.next = new
	ll.len++
	return nil
}

func (ll *SingleLinkedList[T]) RemoveAt(index int) (*T, error) {
	if index > int(ll.len) {
		return nil, fmt.Errorf("Index out of range")
	}

	if index == 0 && ll.Length() == 1 {
		val := ll.head.value
		ll.head = nil
		ll.len = 0
		return &val, nil
	}

	if index == 0 {
		val := ll.head.value
		ll.head = ll.head.next
		ll.len--
		return &val, nil
	}

	curr := ll.head

	for i := 0; i < index-1; i++ {
		curr = curr.next
	}

	value := curr.next.value
	curr.next = curr.next.next
	ll.len--
	return &value, nil

}

func (ll *SingleLinkedList[T]) PrintAll() {
	curr := ll.head
	for curr != nil {
		fmt.Printf("%v, ", curr.value)
		curr = curr.next
	}
	fmt.Println()

}
