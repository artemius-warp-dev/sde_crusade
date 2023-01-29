//User function Template for javascript

/**
 * @param {Node} node
 * @returns {number}
*/

/*
class Node{
    constructor(data){
        this.data = data;
        this.left = null;
        this.right = null;
    }
}
*/

class Solution {
    //Function to find the height of a binary tree.
    height(node)
    {
        //your code here
        if(node == null)
            return 0
        let lh = this.height(node.left)
        let rh = this.height(node.right)
        
        if(lh > rh)
            return lh + 1
        else
            return rh + 1
    }
}
