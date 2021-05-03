package main

import "fmt"
/**

https://leetcode.com/problems/lru-cache/
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists.
Otherwise, add the key-value pair to the cache.
If the number of keys exceeds the capacity from this operation, evict the least recently used key.

Follow up:
Could you do get and put in O(1) time complexity?



Example 1:
Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4

Constraints:

1 <= capacity <= 3000
0 <= key <= 3000
0 <= value <= 104
At most 3 * 104 calls will be made to get and put.
 */

func main()  {
	fmt.Println("initializing test data...")

	fmt.Println("-------------Test case 1---------------")
	cache := NewLRUCache(2)
	cache.Put(1,1)
	cache.Put(2,2)
	fmt.Println("get(1):", cache.Get(1)) //1
	cache.Put(3,3)
	fmt.Println("get(2):", cache.Get(2)) //-1
	cache.Put(4,4)
	fmt.Println("get(1):", cache.Get(1)) //-1
	fmt.Println("get(3):", cache.Get(3)) //3
	fmt.Println("get(4):", cache.Get(4)) //4

	fmt.Println("-------------Test case 2---------------")
	cache2 := NewLRUCache(1)
	cache2.Put(2,1)
	fmt.Println("get(2):", cache2.Get(2)) //1
	cache2.Put(3,2)
	fmt.Println("get(2):", cache2.Get(2)) //-1
	fmt.Println("get(3):", cache2.Get(3)) //2


	fmt.Println("-------------Test case 3---------------")
	cache3 := NewLRUCache(2)
	cache3.Put(2,1)
	cache3.Put(3,2)
	fmt.Println("get(3):", cache3.Get(3)) //2
	fmt.Println("get(2):", cache3.Get(2)) //1
	cache3.Put(4,3)
	fmt.Println("get(2):", cache3.Get(2)) //1
	fmt.Println("get(3):", cache3.Get(3)) //-1
	fmt.Println("get(4):", cache3.Get(4)) //3

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
type LRUCache struct {
	cache map[int]*BiDirectionNode
	capacity int
	head *BiDirectionNode
	tail *BiDirectionNode
}

type BiDirectionNode struct {
	val int
	key int
	pre *BiDirectionNode
	next *BiDirectionNode
}

/**
 Constructor
 */
func NewLRUCache(capacity int) LRUCache {
	return LRUCache{cache: make(map[int]*BiDirectionNode), capacity: capacity}
}


func (this *LRUCache) Get(key int) int {
	node, exists := this.cache[key]
	if !exists {
		return -1
	} else {
		this.moveExistingNodeToHead(node)
		fmt.Println("Head is key = ", this.head.key, "; tail is key = ", this.tail.key)
		return node.val
	}
}


func (this *LRUCache) Put(key int, value int)  {
	node, exists := this.cache[key]
	if exists {
		// update value
		node.val = value
		this.moveExistingNodeToHead(node)
	} else {
		//create new node, put in cache
		newNode := &BiDirectionNode{key: key, val: value}
		this.cache[key] = newNode

		//adjust new tail
		if this.tail == nil {
			this.tail = newNode
		}
		//adjust new head
		if this.head == nil {
			this.head = newNode
		}
		if this.head != newNode {
			this.head.pre = newNode
			newNode.next = this.head
			this.head = newNode
		}

		//evict one if needed
		if len(this.cache) > this.capacity {
			fmt.Println("Evicting cache key = ", this.tail.key)
			delete(this.cache, this.tail.key)
			if this.tail.pre != nil {
				this.tail.pre.next = nil
			}
			this.tail = this.tail.pre
		}
	}
	fmt.Println("Head is key = ", this.head.key, "; tail is key = ", this.tail.key)
}

func (this *LRUCache) moveExistingNodeToHead(node *BiDirectionNode){
	// adjust tail
	if this.tail == node && node.pre != nil {
		this.tail = node.pre
		fmt.Println("Tail now is key = ", this.tail.key)
	}
	// adjust head
	if this.head != node {
		//address existing chain
		if node.next != nil {
			node.next.pre = node.pre
		}
		if node.pre != nil {
			node.pre.next = node.next
		}
		//move to the head
		this.head.pre = node
		node.next = this.head
		this.head = node
		this.head.pre = nil
		fmt.Println("Head now is key = ", this.head.key)
	}
}