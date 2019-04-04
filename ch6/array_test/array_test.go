package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 3, 4, 5}
	t.Log(arr[0], arr[1])
	t.Log(arr1, arr2)
}

func TestArrayTrave(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for _, e := range arr3 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	arr3_sec := arr3[3:3]
	t.Log(arr3_sec)
}
