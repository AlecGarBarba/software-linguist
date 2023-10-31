export function ascending(a: number, b: number): boolean {
  return a > b;
}
export function descending(a: number, b: number): boolean {
  return a < b;
}
export function bubbleSort(
  arr: number[],
  should_sort: (a: number, b: number) => boolean
): void {
  for (let i = 1; i < arr.length; i++) {
    for (let j = 0; j < arr.length - i; j++) {
      const current = arr[j];
      const next = arr[j + 1];
      if (should_sort(current, next)) {
        arr[j] = next;
        arr[j + 1] = current;
      }
    }
  }
}
