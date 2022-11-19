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
 * @param {number} K
 * @return {number}
*/
class Solution {
	// return the Kth largest element in the given BST rooted at 'root'
  	kthLargest(root,K){
  		//code here
  		let counter = 0
  		let res
  		let dfs = function(node) {
  		    if(node == null) return node
  		    dfs(node.right)
  		    counter = counter + 1
  		    if(counter == K)
  		        res = node.data
  		    //console.log([counter, node.data])
  		    dfs(node.left)
  		    
  		}
  		dfs(root)
  		return res
  	}
}
