package bubblesort

type OrderFunction func(a int, b int) bool

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
func BubbleSort(r_arr *[]int, should_order OrderFunction) {
	arr := *r_arr
	array_length := len(arr)
	for i := 1; i < array_length; i++ {
		for j := 0; j < array_length-i; j++ {
			if should_order(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}
