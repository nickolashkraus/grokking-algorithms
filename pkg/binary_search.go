package pkg

// Binary search returns the index of an element in a sorted list in O(log(n))
// time.
//
// If the element is not in the list, return -1.
func binarySearch(elem int, list []int) int {

	if len(list) == 0 {
		return -1
	}

	mid := len(list) / 2

	if list[mid] == elem {
		return mid
	} else if elem < list[mid] {
		// The 'high' value in a[low : high] is excluded from the slice.
		return binarySearch(elem, list[:mid])
	} else if elem > list[mid] {
		return binarySearch(elem, list[mid+1:])
	} else {
		return -1
	}
}
