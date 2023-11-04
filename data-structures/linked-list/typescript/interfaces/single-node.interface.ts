export interface ISingleNode<T> {
  value: T;
  next: ISingleNode<T> | undefined;
}
