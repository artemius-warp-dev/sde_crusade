
/**
 * @param {Node} p
*/

/*
class Node{
    constructor(data){
        this.data = data;
        this.left = null;
        this.right = null;
        this.nextRight = null;
    }
}
*/

class Solution 
{
    //Function to connect nodes at same level.
    connect(p)
    {
        //your code here
        if(p == null) return p
        let queue = []
        let prev = null
        let node = null
        queue.push(p)
        //console.log(queue)
        while(queue.length != 0) {
            let size = queue.length
            for(let i=0; i<size; i++) {
                prev = node
                node = queue.shift()
                //console.log(node)
                if(i > 0) {
                    prev.nextRight = node
                }
                
                if(node.left != null)
                    queue.push(node.left)
                if(node.right != null)
                    queue.push(node.right)
            }
            node.nextRight = null
            
        }
    }
}
