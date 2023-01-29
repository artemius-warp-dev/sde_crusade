//{ Driver Code Starts
//Initial Template for javascript

class Queue{
    constructor(){
        this.arr = [];
        this.front = 0;
    }
    
    push(a){
        this.arr.push(a);
    }
    
    pop(){
        if(this.arr.length != this.front){
            let x = this.arr[this.front];
            this.front++;
            return x;
        }
        else
            return -1;
    }
    
    front_ele(){
        return this.arr[this.front];
    }
    
    empty(){
        if(this.arr.length != this.front)
            return false;
        else
            return true;
    }
}





class QueueStack{
    constructor(){
        this.q1 = new Queue();  
        this.q2 = new Queue();  
    }
    
    /**
     * @param {number} x
    */
    
    //Function to push an element into stack using two queues.
    push(x){
        // code here
        while(!this.q1.empty()) {
            this.q2.push(this.q1.front_ele())
            this.q1.pop()
        }
        this.q1.push(x)
        //console.log([this.q1, this.q2])
        //console.log(this.q2)
        while(!this.q2.empty()) {
            this.q1.push(this.q2.front_ele())
            this.q2.pop()
        }
        
    }
    
    /**
     * @returns {number}
    */
    
    //Function to pop an element from stack using two queues. 
    pop(){
        // code here
        //console.log([this.q1, this.q2])
        return this.q1.pop()
        
    }
}

//THis is more interested variant
    // push(x){

    //     if(this.q1.arr.length==0){this.q1.push(x)}

    //     else{

    //         this.q2.push(x)

    //         while(this.q1.arr.length>this.q1.front){

    //             this.q2.push(this.q1.pop())

    //         }

    //         this.q1=this.q2

    //         this.q2=new Queue

    //     }

        

    // }

    // pop(){

    //     return this.q1.pop()
    // }


let s = new QueueStack()

// q1:
// 2 3, 2, 4, 2

// q2:

// 2 2

// 3



    
s.push(2)
s.push(3)
console.log(s.pop())
s.push(4)
console.log(s.pop())
