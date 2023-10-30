package bubblesort

type orderFunction func(a int, b int) bool

func Ascending(a int, b int) bool {
	return (a > b)
}

func Descending(a int, b int) bool {
	return a < b
}

/*
* Simple Bubblesort algorithm, takes an array of integers,
* returns it ordered based on an orderFunction.
 */
func BubbleSort(arr *[]int, should_order orderFunction) {

	for i := 1; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i; j++ {
			if should_order((*arr)[j], (*arr)[j+1]) {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}

}
