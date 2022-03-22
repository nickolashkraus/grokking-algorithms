package pkg

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 3, 5, 7, 9}
	ret := binarySearch(3, list)
	if ret != 1 {
		t.Fail()
	}
	ret = binarySearch(4, list)
	if ret != -1 {
		t.Fail()
	}
}
