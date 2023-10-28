import { binary_search } from "./binary_search";

function createOrderedArray(n: number): number[] {
  return Array.from(Array(n).keys());
}

function getPerformanceTime(performanceFunc: () => void): number {
  const t0 = performance.now();

  performanceFunc();

  const t1 = performance.now();

  return t1 - t0;
}

function getAverage(arr: number[]): number {
  return arr.reduce((prev, curr) => prev + curr, 0) / arr.length;
}

function doPerformanceBenchmark(array_size: number): void {
  const array = createOrderedArray(array_size);

  const results: number[] = [];

  for (let i = 0; i < array_size; i++) {
    results.push(getPerformanceTime(() => binary_search(array, array[i])));
  }

  console.log(
    `Average time for array of length ${array_size}: ${
      getAverage(results) * 1000
    } us`
  );
}
/**
 * First Scenario: binary search for 10_000 elements
 + And do a binary seaarch looking at every element
 */

doPerformanceBenchmark(10_000);

/**
 * Second scenario: binary search for 100_000 elements,
 * we are doing this 500 times, saving results in json
 */

doPerformanceBenchmark(100_000);

doPerformanceBenchmark(1_000_000);

doPerformanceBenchmark(10_000_000);
doPerformanceBenchmark(100_000_000);
