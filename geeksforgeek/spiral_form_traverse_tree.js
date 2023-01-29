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
    //Function to return a list containing the level order 
	//traversal in spiral form.	
    findSpiral(root)
    {
        //your code here
        let queue_nodes = []
        let result = []
        if(root == null)
            return result
        queue_nodes.push(root)
        let left_to_right = false
        while(queue_nodes.length != 0) {
            let row = []
            let size = queue_nodes.length

            for(let i=0; i<size; i++) {
                let node = queue_nodes.shift()
                let index = left_to_right ? i : size - 1 - i
                
                row[index] = node.data
                if(node.left != null)
                    queue_nodes.push(node.left)
                if(node.right != null)
                    queue_nodes.push(node.right)
            }
            left_to_right = !left_to_right
            for(let i=0; i< row.length; i++)
                result.push(row[i])
            
        }
        return result
        
    }
}
