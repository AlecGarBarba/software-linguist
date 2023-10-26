import { binary_search } from "./binary_search";

describe("Binary Search", () => {
  it("Should return false if the array is empty", () => {
    const empty_array: number[] = [];

    expect(binary_search(empty_array, 1)).toBeFalsy();
  });

  it("Return false if the element does not exist in the array", () => {
    const even_number_array: number[] = [2, 4, 6, 8, 10, 12, 14, 16];

    expect(binary_search(even_number_array, 1)).toBeFalsy();
  });

  it("Should return true if element is at the start of the array", () => {
    const simple_array: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

    expect(binary_search(simple_array, simple_array[0]));
  });

  it("Should return true if element is at the start of the array", () => {
    const simple_array: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

    expect(binary_search(simple_array, simple_array[0])).toBeTruthy();
  });

  it("Should return true if element is at the middle of the array", () => {
    const simple_array: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

    expect(
      binary_search(
        simple_array,
        simple_array[Math.round(simple_array.length / 2)]
      )
    ).toBeTruthy();
  });
  it("Should return true if element is at the end of the array", () => {
    const simple_array: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

    expect(
      binary_search(simple_array, simple_array[simple_array.length - 1])
    ).toBeTruthy();
  });
});
