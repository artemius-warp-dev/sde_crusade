// User function Template for javascript

/**
 * @param {Node} root
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
    // Function to return the diameter of a Binary Tree.
    diameter(root) {
        // your code here
        let diameter = [0]
        let height = function(node, d) {
            if(node == null)
                return 0
            let lh = height(node.left, d)
            let rh = height(node.right, d)
            //console.log([lh, rh, d[0]])
            d[0] = Math.max(d[0], lh + rh + 1)
            return 1 + Math.max(lh, rh)
        }
        
        height(root, diameter)
        return diameter[0]
    }
}
