//User function Template for javascript

class Node {
    constructor(x) {
        this.data = x
        this.next = null

    }
}


class MyQueue{
    constructor(){
        this.front = null;  // QueueNode
        this.rear = null;   // QueueNode
    }
    
    /**
     * @param {number} x
    */
    //Function to push an element into the queue.
    push(x){
        // code here
        
        let node = new Node(x)
        if(this.front == null) {
            this.front = node
            this.rear = this.front
        } else {
            this.rear.next = node
            this.rear = this.rear.next
        }
            
    }
    
    /**
     * @returns {number}
    */
    //Function to pop front element from the queue.
    pop(){
        // code here
       
        if(this.front != null) {
            let value = this.front.data
            this.front = this.front.next
            return value
            
        } else {
            return -1
        }
            
        
    }
}

let s = new MyQueue()
s.push(2)
s.push(3)
console.log(s.pop())
s.push(4)
console.log(s.pop())
console.log(s.front)
