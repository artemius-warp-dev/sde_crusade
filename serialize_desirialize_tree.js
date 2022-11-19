/*
class Node{
    constructor(data){
        this.data = data;
        this.left = null;
        this.right = null;
    }
}
*/

class Solution 
{
    /**
     * @param {Node} root
     * @returns {number[]}
    */
    
    //Function to serialize a tree and return a list containing nodes of tree.
    serialize(root)
    {
        //your code here
        let res = []
        let q = []
        if(root == null) return res
        q.push(root)
        while(q.length != 0) {
            let node = q.shift()
            if(node == null) {
                res.push('N')
            } else {
                res.push(node.data)
                q.push(node.left)
                q.push(node.right)
            }
        }
        //console.log(res)
        return res
    }
    
    /**
     * @param {number[]} A
     * @returns {Node}
    */
    
    //Function to deserialize a list and construct the tree
    deSerialize(A)
    {
        //your code here
        if(A.length == 0) return null
        let q = []
        let root = new Node(A.shift())
        q.push(root)
        
        while(q.length != 0) {
            let node = q.shift()
            
            let e = A.shift()
            
            if(e == 'N') {
                node.left = null
            } else {
                let node_left = new Node(e)
                node.left = node_left
                q.push(node_left)
            }
            
            e = A.shift()
            if(e == 'N') {
                node.right = null
            } else {
                let node_right = new Node(e)
                node.right = node_right
                q.push(node_right)
            }
        }    
        return root
        
    }
}
