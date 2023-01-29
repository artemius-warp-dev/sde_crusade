




class Node {
    constructor() {
        this.data = ""
        this.prev = null
        this.next = null
    }
}

class Deque {
    constructor() {
        this.head = null
        this.tail = this.head
    }
    
    append_node(x) {
        let temp = new Node()
        temp.data = x
        temp.prev = temp.next = null

        if(this.head == null) {
            this.head = this.tail = temp
            return 
        }

        this.tail.next = temp
        temp.prev = this.tail
        this.tail = temp

        
    }

    remove_node(node) {
        if(this.head == node)
            this.head = this.head.next
        if(this.tail == node) {
            this.tail = this.tail.prev
            
        }
        if(node.next != null)
            node.next.prev = node.prev
        if(node.prev != null)
            node.prev.next = node.next

    }
}

class Solution {
    constructor() {
        this.deque = new Deque()
    }
    
    findFirstNonRepeating(stream) {
        let inDLL = []
        let repeated = []
        let MAX_CHAR = 256
        let res_str = ""
        for(let i=0; i < MAX_CHAR; i++) {
            inDLL[i] = null
            repeated[i] = false
        }

        for(let i=0; stream[i]; i++) {
            let ch = stream[i]
            if(!repeated[ch]) {
                if(inDLL[ch] == null) {
                    this.deque.append_node(ch)
                    inDLL[ch] = this.deque.tail
                } else {
                    this.deque.remove_node(inDLL[ch])
                    inDLL[ch] = null
                    repeated[ch] = true
                 }
            }

            if(this.deque.head != null)
                res_str += this.deque.head.data
            else
                res_str += '#'
        }
        
        return res_str
        
    }
        
}

let stream = "zz"
let s = new Solution()
console.log(s.findFirstNonRepeating(stream))



//d = new Deque()
// d.append_node(10)
// d.append_node(3)
// d.append_node(5)
// console.log(d)
// d.remove_node(d.head.next)
// console.log(d)
