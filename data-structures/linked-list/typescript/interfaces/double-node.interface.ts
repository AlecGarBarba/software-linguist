interface IDoubleNode<T> {
  value: T;
  previous: IDoubleNode<T> | null;
  next: IDoubleNode<T> | null;
}
