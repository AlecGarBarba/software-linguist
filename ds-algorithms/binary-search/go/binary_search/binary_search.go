package binarysearch

func Binarysearch(arr *[]uint32, search uint32) bool {

	high := len(*arr)

	if high == 0 {
		return false
	}
	low := 0

	for low < high {

		middle_point := low + (high-low)/2

		value := (*arr)[middle_point]

		if value == search {
			return true
		}
		// if the value obtained is greater than what I'm looking for
		// this means that we need to search in the lower section of the array
		if value > search {
			high = middle_point
			//else, we want to search in the upper section of the array
		} else {
			low = middle_point + 1
		}

	}
	return false
}
