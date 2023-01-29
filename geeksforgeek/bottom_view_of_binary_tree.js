
//User function Template for javascript

/**
 * @param {Node} root
 * @returns {number[]}
*/

class Pair {
    constructor(first, second) {
        this.first = first
        this.second = second
    }
}

class Solution
{
    //Function to return a list containing the bottom view of the given tree.
    bottomView(root)
    {
        //your code here
        let result = []
        if(root == null) return result
        let map = new Map()
        let queue_nodes = []
        queue_nodes.push(new Pair(root, 0))
        
        while(queue_nodes.length != 0) {
            let node = queue_nodes.shift()
            map.set(node.second, node.first.data)
            
            if(node.first.left != null)
                queue_nodes.push(new Pair(node.first.left, node.second - 1))
            if(node.first.right != null)
                queue_nodes.push(new Pair(node.first.right, node.second + 1))
        }
        
        let keys = Array.from(map.keys()).sort((a,b) => a-b)
        //console.log(map)
        
        for(let i = 0; i<keys.length; i++) {
            result.push(map.get(keys[i]))
        }
            
        return result
        
    }
}
