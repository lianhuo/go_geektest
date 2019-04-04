package map_ext_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[0] = func(op int) int { return op }
	m[1] = func(op int) int { return op * op }
	m[2] = func(op int) int { return op * op * op }
	t.Log(m[0](2), m[1](2), m[2](2))
	t.Log(m)
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[2] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	t.Log(len(mySet))
	// delete(mySet, 1)
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	t.Log(len(mySet))

}
