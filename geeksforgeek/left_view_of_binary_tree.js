//User function Template for javascript

/**
 * @param {Node} root
 * @returns {number[]}
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
    leftView(root)
    {
        //your code here
        let cache = []
        let _left_view = function(root, level) {
            if(root == null) return
            
            if(level == cache.length)
                cache.push(root.data)
            
            _left_view(root.left, level + 1)
            _left_view(root.right, level + 1)
            
        }
        _left_view(root, 0)
        
        return cache
        
    }
    
    
    
}
