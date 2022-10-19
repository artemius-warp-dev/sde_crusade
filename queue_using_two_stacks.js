//User function Template for javascript


class Stack{
    constructor(){
        this.arr = [];
        this.top = -1;
    }
    
    push(a){
        this.arr.push(a);
        this.top++;
    }
    
    pop(){
        this.arr.pop();
        this.top--;
    }
    
    front(){
        return this.arr[this.top];
    }
    
    empty(){
        if(this.top != -1)
            return false;
        else
            return true;
    }
}




class StackQueue{
    constructor(){
        this.s1 = new Stack();
        this.s2 = new Stack();
    }
    
    /**
     * @param {number} B
     */

    // 5
    // 1 2 1 3 2 1 4 2

    // s1:
    //4

    // s2:
    //3
    //2 pop
    
    //Function to push an element in queue by using 2 stacks.
    push(B){
        // code here
        this.s1.push(B)
    }
    
    /**
     * @returns {number}
    */
    
    //Function to pop an element from queue by using 2 stacks.
    pop(){
        // code here
       // console.log(this.s1.top)
        if(this.s1.empty() && this.s2.empty())
            return -1
        
        while(!this.s1.empty()) {
            this.s2.push(this.s1.front())
            this.s1.pop()
        }
        let elem = this.s2.front()
        this.s2.pop()

        while(!this.s2.empty()) {
            this.s1.push(this.s2.front())
            this.s2.pop()
        }
        
        return elem
    }
}

let s = new StackQueue()
s.push(2)
s.push(3)
console.log(s.pop())
s.push(4)
console.log(s.pop())


