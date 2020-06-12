package main

import (
	"math/rand"
)

type RandomizedSet struct {
	values    map[int]int
	positions []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		values:    make(map[int]int),
		positions: make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (s *RandomizedSet) Insert(val int) bool {
	_, ok := s.values[val]
	if ok {
		return false
	}

	s.values[val] = len(s.positions)
	s.positions = append(s.positions, val)

	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (s *RandomizedSet) Remove(val int) bool {
	position, ok := s.values[val]
	if !ok {
		return false
	}
	delete(s.values, val)

	last := len(s.positions) - 1
	lastVal := s.positions[last]
	if val != lastVal {
		s.positions[position] = lastVal
		s.values[lastVal] = position
	}
	s.positions = s.positions[:last]

	return true
}

/** Get a random element from the set. */
func (s *RandomizedSet) GetRandom() int {
	length := len(s.positions)
	return s.positions[rand.Intn(length)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
