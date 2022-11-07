class StackNode{
    constructor(a){
        this.data = a;
        this.next = null;
    }
}



//User function Template for javascript


class MyStack
{
    
    constructor(){
        this.top = null
        
    }
    
    /**
     * @param {number} a
    */
    
    push(a){
        // code here
        let node = new StackNode(a)
        if(this.top == null) {
            this.top = node
        } else {
            node.next = this.top
            this.top = node
            //this.top = this.top.next
        }
    }
    
    /**
     * @returns {number}
    */
    //Function to pop front element from the queue.
    pop(){
        // code here
       
        if(this.top != null) {
            let value = this.top.data
            
            this.top = this.top.next
            return value
            
        } else {
            return -1
        }
            
        
    }
}


let s = new MyStack()
s.push(2)
s.push(3)
console.log(s.pop())
s.push(4)
console.log(s.pop())
console.log(s.top)
