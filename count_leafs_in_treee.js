//User function Template for javascript

/*
class Node
{
    constructor(x){
        this.data=x;
        this.left=null;
        this.right=null;
    }
}
*/

/**
 * @param {Node} root
 * @return {number} 
*/
/* Should return count of leaves. For example, return
    value should be 2 for following tree.
         10
      /      \ 
   20       30 */

class Solution {
  	countLeaves(root){
  		//code here
  		if(root == null) return 0
  		if(root.left == null && root.right == null) return 1
  		return this.countLeaves(root.left) + this.countLeaves(root.right) 
  	}
}
