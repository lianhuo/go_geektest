package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m1)
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[2] = 3
	t.Log(m2)
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Log(m3)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[0])
	m1[1] = 0
	t.Log(m1[1])
	m1[2] = 1
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("key 3 is existing value=%d", v)
	} else {
		t.Log("key 3 is not existing")
	}
}

func TestTraveMap(t *testing.T){
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	for k,v:=range m1{
		t.Log(v,k)
	}
}