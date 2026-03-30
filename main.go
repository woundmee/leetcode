package main

import "fmt"

// problem -> https://neetcode.io/problems/design-hashmap/
// задача: реализовать кастомные методы для своего хеш-мапы

func main() {

	myHashMap := MyHashMap{
		dict: [][]int{},
	}

	myHashMap.Put(3, 4)
	myHashMap.Put(5, 3)
	myHashMap.Put(2, 1)
	myHashMap.Put(5, 19)

	fmt.Println(myHashMap)

	// myHashMap.Put(1, 1) // The map is now [[1,1]]
	// myHashMap.Put(2, 2) // The map is now [[1,1], [2,2]]
	// myHashMap.Get(1)    // return 1, The map is now [[1,1], [2,2]]
	// myHashMap.Get(3)    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
	// myHashMap.Put(2, 1) // The map is now [[1,1], [2,1]] (i.e., update the existing value)
	// myHashMap.Get(2)    // return 1, The map is now [[1,1], [2,1]]
	// myHashMap.Remove(2) // remove the mapping for 2, The map is now [[1,1]]
	// myHashMap.Get(2)    // return -1 (i.e., not found), The map is now [[1,1]]

	// fmt.Println("myHashMap =", myHashMap)

}

type MyHashMap struct {
	dict [][]int // note: можно было юзать просто map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		dict: [][]int{},
	}
}

func (this *MyHashMap) Put(key int, value int) {

	// 1. добавление в слайс
	// 2. изначально пустой слайс
	// 3. если ключ уже есть в слайсе, обновить его
	// 4. если ключа нет - добавить в слайс

	for i := range this.dict {
		if this.dict[i][0] == key {
			this.dict[i] = []int{key, value}
			return
		}
	}

	// ключа нет - добавляем
	this.dict = append(this.dict, []int{key, value})

}

func (this *MyHashMap) Get(key int) int { // ok

	for i := range this.dict {
		if key == this.dict[i][0] {
			return this.dict[i][1]
		}
	}

	return -1
}

// note: можно было просто вызвать delete()
func (this *MyHashMap) Remove(key int) { // ok
	// key=5
	// [ [1,2], [3,4], [5,2], [4,1] ]
	//                   ⤷ remove

	if len(this.dict) < 1 {
		return
	}

	for i := range this.dict {
		if this.dict[i][0] == key {
			this.dict = append(this.dict[:i], this.dict[i+1:]...) // удалил подмассив
			return
		}
	}

}
