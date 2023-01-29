//User function Template for javascript

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

/**
 * @param {Node} root
 * @return {boolean}
*/
class Solution {
  	// return true/false denoting whether the tree is Symmetric or not
    isSymmetric(root){
        //code here
        return (root == null || this.symetric_traverse(root.left, root.right))
    }
    
    symetric_traverse(left, right) {
        if(left == null || right == null) 
            return (left == right)
        return (left.key == right.key) && this.symetric_traverse(left.left, right.right) && this.symetric_traverse(left.right, right.left) 
    }
}
