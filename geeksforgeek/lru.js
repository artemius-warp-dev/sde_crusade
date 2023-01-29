class Node {
    constructor(key,value) {
        this.key = key
        this.value = value
        this.prev = null
        this.next = null
        
    }
}

class Deque {
    constructor() {
        this.head = new Node(-1,-1)
        this.tail = new Node(-1,-1)
        this.head.next = this.tail
        this.tail.prev = this.head
        
    }
    
    insert(node) {
        let head_next = this.head.next
        this.head.next = node
        node.prev = this.head
        head_next.prev = node
        node.next = head_next
    }

    remove(node) {
        node.prev.next = node.next
        node.next.prev = node.prev
    }    
 

}


//User function Template for javascript

class LRUCache 
{
    //Constructor for initializing the cache capacity with the given value.
    constructor(cap)
    {
        //your code here
        this.d = new Deque()
        this.cache = new Map()
        this.capacity = cap
        
    }
    
    
    /**
     * @param {number} key
     * @returns {number}
    */
    
    //Function to return value corresponding to the key.
    get(key)
    {
        //your code here
        if(this.cache.has(key)) {
            let node = this.cache.get(key)
           
            this.cache.delete(node.key)
            this.d.remove(node)
         
            this.cache.set(key,  node)
            
            this.d.insert(node)
            return node.value
        } else
            return -1
        
    }

    
    /**
     * @param {number} key
     * @param {number} value
    */
    
    //Function for storing key-value pair.
    set(key, value)
    {
        //your code here
        let node = new Node(key, value)
        
        if(this.cache.has(key)) {
            let node = this.cache.get(key)
            this.cache.delete(node.key)
            this.d.remove(node)
        }
         
        if(this.cache.size == this.capacity) {
            this.cache.delete(this.d.tail.prev.key)
            this.d.remove(this.d.tail.prev)
        }

        this.cache.set(key, node) 
        this.d.insert(node)
    }
}


let s = new LRUCache()
s.set(1,2)
s.set(2,3)
s.set(1,5)
//console.log(s.d)
console.log(s.get(2))
console.log(s.d)
