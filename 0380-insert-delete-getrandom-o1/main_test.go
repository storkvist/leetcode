package main

import "testing"

func TestRandomizedSet(t *testing.T) {
	obj := Constructor()
	ok := obj.Insert(1)
	if !ok {
		t.Errorf("%v.Insert(1) = false; want true", obj)
	}
	ok = obj.Insert(1)
	if ok {
		t.Errorf("%v.Insert(1) = true; want false", obj)
	}
	obj.Insert(2)
	obj.Insert(3)
	obj.Insert(4)
	ok = obj.Remove(4)
	if !ok {
		t.Errorf("%v.Remove(4) = false; want true", obj)
	}
	ok = obj.Remove(4)
	if ok {
		t.Errorf("%v.Remove(4) = true; want false", obj)
	}
	obj.GetRandom()
}
