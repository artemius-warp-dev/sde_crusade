//User function Template for javascript

/**
 * @param {Node} root
 * @returns {Node}
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
    //Function to convert a binary tree to doubly linked list and return it.
    bToDLL(root)
    {
        //your code here
        
        
        let in_order_traverse = function(root) {
            if(root == null) return
            
            let left = root.left
            let right = root.right
            
            in_order_traverse(root.left)
            
            if(head == null) {
                head = root
                //console.log(head)
            }
                
            root.left = tail
            
            if(tail != null) {
                tail.right = root
            }
            
            tail = root
            
            in_order_traverse(root.right)
            
        }
        
        
        let head = null
        let tail = null
        
        if(root == null) {
            head = root
        }
        
        in_order_traverse(root)
        
        return head
        
        
    }
}
