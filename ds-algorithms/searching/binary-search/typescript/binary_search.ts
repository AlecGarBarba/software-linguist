export const binary_search = (arr: number[], expected: number): boolean => {
  if (arr.length === 0) {
    return false;
  }

  let low: number = 0;

  let high: number = arr.length;

  let value: number;

  do {
    let midpoint: number = Math.floor(low + (high - low) / 2);
    value = arr[midpoint];

    if (value === expected) {
      return true;
    }

    // if the value obtained is greater than what I'm looking for
    // this means that we need to search in the lower section of the array
    if (value > expected) {
      high = midpoint;
      //else, we want to search in the upper section of the array
    } else {
      low = midpoint + 1;
    }
  } while (low < high);

  return false;
};
