package sorting



func Bubble_sort_modified(list []int) {
	n := len(list)
	swapped := true
	for swapped {
		swapped = false
		for x := 1; x < n-1; x++ {
			if list[x-1] > list[x] {
				// swap values using Go's tuple assignment
				list[x], list[x-1] = list[x-1], list[x]
				swapped = true
			}
		}
	}


}

// Implementering av Bubble_sort algoritmen
func Bubble_sort(list []int) {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
}

// Implementering av Quicksort algoritmen
func QSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}
