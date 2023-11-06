Instruction:

Create a Single linked list Data structure following this interface

```ts
interface LinkedList<T> {
  get length(): number;
  insertAt(item: T, index: number): void;
  remove(item: T): T | undefined;
  removeAt(index: number): T | undefined;
  append(item: T): void;
  prepend(item: T): void;
  get(index: number): T | undefined;
}
```

```go

type ILinkedList[T interface{}] interface {
	Length() uint
	InsertAt(item T, index uint)
	RemoveAt(index uint) (T, error)
	Append(item T)
	Prepend(item T)
	Get(index uint) (T, error)
}
```
