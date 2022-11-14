//User function Template for javascript

/**
 * @param {Node} root
 * @returns {boolean}
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
    
    //Function to check whether a binary tree is balanced or not.
    isBalanced(root)
    {
        //your code here
        return this.dfs_height(root) != -1
    }
    
    dfs_height(root) {
        if(root == null) return 0
        let lh = this.dfs_height(root.left)
        if(lh == -1) return -1
        let rh = this.dfs_height(root.right)
        if(rh == -1) return -1
        if(Math.abs(lh - rh) > 1) return -1 
        return Math.max(lh, rh) + 1
    }
}
