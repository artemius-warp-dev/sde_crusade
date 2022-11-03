//User function Template for javascript

class MyStack {
    constructor() {
        this.stack = []
        this.min = 0
  }
  
  /**
   * @param {number} x
   */
  
  /* The method push to push element into the stack */
  push(x) {
      // code here
      if (this.stack.length == 0) {
          this.min = x
          this.stack.push(x)
      } else if(x >= this.min) {
          this.stack.push(x)
      } else {
          let diff = 2 * x - this.min
          this.stack.push(diff)
          this.min = x
      }
      
  }

  /**
   * @returns {number}
   */
  
  /*The method pop which return the element poped out of the stack*/
  pop() {
      // code here

     // console.log(this.stack)
      
      let top = function(stack) {
          return stack[stack.length  - 1]
      }

      if(this.stack.length == 0)
          return -1

      if(top(this.stack) > this.min)
          return this.stack.pop()
      else {
          console.log(this.stack)
          let diff = 2 * this.min - top(this.stack)
          let e = this.min
          this.min = diff
          this.stack.pop()
          return e
      }
  }
  
  /**
   * @returns {number}
   */
   
  /*The method getMin() which return the minimum element of the stack*/
  getMin() {
      // code here
      //console.log([this.min, this.stack])
      // console.log(this.min)
      if(this.stack.length == 0)
          return -1
      else
          return this.min
  }
}

let s = new MyStack()
s.push(79)
s.getMin()
s.push(4)
s.getMin()
s.getMin()
console.log(s.pop())
s.push(61)
s.getMin()
