package pkg

// A node in a linked list.
type Node struct {
	val  int
	next *Node
}

// Selection sort divides the input list into two parts: a sorted sublist of
// items which is built up from left to right at the front (left) of the list
// and a sublist of the remaining unsorted items that occupy the rest of the
// list.
//
// Initially, the sorted sublist is empty and the unsorted sublist is the
// entire input list. The algorithm proceeds by finding the smallest element in
// the unsorted sublist, swapping it with the leftmost unsorted element, and
// moving the sublist boundaries one element to the right.
func selectionSortArray(list []int) []int {
	for i := 0; i < len(list); i++ {
		curr, smallest := i, i
		for i := curr; i < len(list); i++ {
			if list[i] < list[smallest] {
				smallest = i
			}
		}
		temp := list[i]
		list[i] = list[smallest]
		list[smallest] = temp
	}
	return list
}
