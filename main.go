package main

import "fmt"

// problem -> https://neetcode.io/problems/design-hashset/
// задача: релазиовать свои хеш-функции
func main() {
	m := MyHashSet{m: make(map[int]struct{})}
	m.Add(1)
	m.Add(2)
	m.Add(3)

	fmt.Println(m)

	m.Remove(3)

	fmt.Println(m)
}

type MyHashSet struct {
	m map[int]struct{}
}

func Constructor() MyHashSet {
	return MyHashSet{
		m: make(map[int]struct{}),
	}
}

func (this *MyHashSet) Add(key int) {
	if _, ok := this.m[key]; ok {
		return
	}
	this.m[key] = struct{}{} // добавляем новый ключ
}

func (this *MyHashSet) Remove(key int) {
	delete(this.m, key)
}

func (this *MyHashSet) Contains(key int) bool {
	if _, ok := this.m[key]; ok {
		return true
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
