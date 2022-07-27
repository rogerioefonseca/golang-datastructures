package main

import "fmt"

//Array size of the hashtable
const arraySize = 7

var ht hashTable

type hashTable struct {
	array [arraySize]*bucket
}

func (h *hashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

//func (h *hashTable) delete(key string) {
//	index := hash(key)
//}

func hash(key string) int {
	sum := 0
	for i, _ := range key {
		sum += int(key[i])
	}

	return sum % arraySize
}

type bucket struct {
	head *bucketNode
}

func (b *bucket) insert(k string) {
	if b.search(k) {
		return
	}
	newNode := bucketNode{key: k}
	newNode.next = b.head
	b.head = &newNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func init() {
	//node1 := bucketNode{key: "rogerio"}
	//b := bucket{ head: }
	for i, _ := range ht.array {
		ht.array[i] = &bucket{}
	}
}

func main() {
	fmt.Println(ht)
	fmt.Println(hash("RANDY"))
	testBucket := &bucket{}
	testBucket.insert("test")
	fmt.Println(testBucket.search("test"))
	fmt.Println(testBucket.search("test1"))
}
