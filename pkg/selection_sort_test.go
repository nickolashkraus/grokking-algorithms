package pkg

import (
	"reflect"
	"testing"
)

func TestSelectionSortArray(t *testing.T) {
	list := []int{5, 3, 6, 2, 10}
	ret := selectionSortArray(list)
	if !reflect.DeepEqual(ret, []int{2, 3, 5, 6, 10}) {
		t.Fail()
	}
}
