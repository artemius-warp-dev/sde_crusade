
/*
class Node
{
    constructor(x){
        this.key=x;
        this.left=null;
        this.right=null;
    }
}
*/

class NodeValue {
    constructor(min, max, size) {
        this.size = size
        this.min = min
        this.max = max
    }
}

/**
 * @param {Node} root
 * @return {number}
*/
class Solution {
  	largestBst(root){
  		//code here
  		
  		
  		let f = function(root) {
  		    if(root == null) return new NodeValue(Number.MAX_VALUE, Number.MIN_VALUE, 0)
  		    
  		    let left = f(root.left)
  		    let right = f(root.right)
  		    //console.log([left.max, root.key, right.min])
  		    if(left.max < root.key && root.key < right.min) {
  		        return new NodeValue(Math.min(root.key, left.min), Math.max(root.key, right.max), left.size + right.size + 1)
  		    }
  		    
  		    return new NodeValue(Number.MIN_VALUE, Number.MAX_VALUE, Math.max(left.size, right.size))
  		}
  		return f(root).size
  	}
}
