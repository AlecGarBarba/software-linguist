import { ascending, bubbleSort, descending } from ".";

describe("Test", () => {
  it("Should order ascending", () => {
    const arr = [1, 3, 7, 2, 4, 6, 5, 8];

    bubbleSort(arr, ascending);

    expect(arr).toEqual([1, 2, 3, 4, 5, 6, 7, 8]);
  });
});

describe("Test", () => {
  it("Should order descending", () => {
    const arr = [1, 3, 7, 2, 4, 6, 5, 8];

    bubbleSort(arr, descending);

    expect(arr).toEqual([8, 7, 6, 5, 4, 3, 2, 1]);
  });
});

describe("Test", () => {
  it("Should do nothing on empty", () => {
    const arr = [];
    bubbleSort(arr, ascending);

    expect(arr).toEqual([]);
  });
});
