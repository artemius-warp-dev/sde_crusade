
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

class QueueNode {
    constructor(vertical, level, node) {
        this.vertical = vertical
        this.level = level
        this.node = node
    }
}

class Level {
    constructor(l) {
        this.nodes = []
        this.l = l
    }
    
    add(data) {
        this.nodes.push(data)
    }
}

class Verticle {
    constructor(v) {
        this.levels = new Map()
        this.v = v
    }
}

class Cache {
    constructor() {
        this.cache = new Map()
        this.levels = new Map()
        this.nodes = []
    }
    add(v, l, n) {
        let level
        let verticle
        
        if(!this.cache.has(v))
            this.cache.set(v, new Verticle(v))
        
        verticle = this.cache.get(v)
        
        if(!verticle.levels.has(l))
            verticle.levels.set(l, new Level(l))
            
        level = verticle.levels.get(l)
        level.add(n)
        
        
        
    
        verticle.levels.set(l, level)
        this.cache.set(v, verticle)
      
        
    }
    
}

class Solution 
{
    //Function to find the vertical order traversal of Binary Tree.
    verticalOrder(root)
    {
        //your code here
        let result = []
        if(root == null) return result
        
        let queue = []
        queue.push(new QueueNode(0, 0, root))
        let cache = new Cache()
        while(queue.length != 0) {
            let qnode = queue.shift()
            //console.log(qnode)
            cache.add(qnode.vertical, qnode.level, qnode.node.data)
            if(qnode.node.left != null) 
                queue.push(new QueueNode(qnode.vertical - 1, qnode.level + 1, qnode.node.left))
            if(qnode.node.right != null)
                queue.push(new QueueNode(qnode.vertical + 1, qnode.level + 1, qnode.node.right))
            
        } 
        
        //console.log(cache)
        let sorted_cache = new Map([...cache.cache.entries()].sort((a,b) => Number(a[0]) -Number(b[0])  ))
        //console.log(sorted_cache)
        for (const [vkey, vvalues] of sorted_cache) {
            for(const [lkey, lvalues] of vvalues.levels) {
                //console.log([vkey, lvalues.nodes])
                result = result.concat(lvalues.nodes)
            }
        }
        return result
        
    }
}
