#Using Doubly Linked List and Hash Map
type LRUCache struct {
    capacity int
    cache    map[int]*DLinkedNode
    head     *DLinkedNode
    tail     *DLinkedNode
}

type DLinkedNode struct {
    key, value int
    prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
    cache := LRUCache{
        capacity: capacity,
        cache:    make(map[int]*DLinkedNode),
        head:     &DLinkedNode{},
        tail:     &DLinkedNode{},
    }
    cache.head.next = cache.tail
    cache.tail.prev = cache.head
    return cache
}

func (this *LRUCache) Get(key int) int {
    if node, exists := this.cache[key]; exists {
        this.moveToHead(node)
        return node.value
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exists := this.cache[key]; exists {
        node.value = value
        this.moveToHead(node)
    } else {
        node := &DLinkedNode{key: key, value: value}
        this.cache[key] = node
        this.addNode(node)
        if len(this.cache) > this.capacity {
            tail := this.popTail()
            delete(this.cache, tail.key)
        }
    }
}

func (this *LRUCache) addNode(node *DLinkedNode) {
    node.prev = this.head
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache)
