//User function Template for javascript

/**
 * @param {Node} root
 * @returns {boolean}
*/


class Solution 
{
    //Function to check whether a Binary Tree is BST or not.
    isBST(root)
    {
        //your code here
        if(root == null) return null
        let min = Number.MIN_VALUE
        let max = Number.MAX_VALUE
        let is_bst = function(node, min_bound, max_bound) {
            if(node == null) return true
            if(node.data >= max_bound || node.data <= min_bound) return false
            return is_bst(node.left, min_bound, node.data) &&
            is_bst(node.right, node.data, max_bound)
        }
        
        return is_bst(root, min, max)
    }
}
