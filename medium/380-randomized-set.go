package main

import (
	"math/rand"
)

// Idea: set for inserting/deleting, randarray for random values
// Vals are keys in set, values in set are indices of corresponding keys in randarray
type RandomizedSet struct {
	set       map[int]int
	randarray []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set:       make(map[int]int),
		randarray: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok != true {
		this.set[val] = len(this.randarray)
		this.randarray = append(this.randarray, val)
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.set[val]; ok == true {
		temp := this.randarray[len(this.randarray)-1]
		this.randarray[len(this.randarray)-1] = this.randarray[this.set[val]]
		this.randarray[this.set[val]] = temp
		this.set[temp] = this.set[val]

		this.randarray = this.randarray[:len(this.randarray)-1]

		delete(this.set, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.randarray[rand.Intn(len(this.randarray))]
}
