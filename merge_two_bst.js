//User function Template for javascript

/**
 * @param {Node} root1
 * @param {Node} root2
 * @returns {number[]}
*/



class Solution
{
    //Function to return a list of integers denoting the node 
    //values of both the BST in a sorted order.
    merge(root1, root2)
    {
        //your code here
        let list1 = this.store_in_order(root1)
        let list2 = this.store_in_order(root2)
        let merged_list = this.merge_lists(list1, list2)
        //console.log(merged_list)
        //console.log( this.build_bst(merged_list, 0, merged_list.length  - 1))
        let merged_bst = this.build_bst(merged_list, 0, merged_list.length  - 1)
        return this.store_in_order(merged_bst)
        
    }
    
    build_bst(list, start, end) {
        if(start > end) {
            return null
        }
        
        let mid = parseInt((start + end) / 2)
        let node = new Node(list[mid])
        node.left = this.build_bst(list, start, mid - 1)
        node.right = this.build_bst(list, mid + 1, end)
        
        return node
        
    }
    
    merge_lists(list1, list2) {
        let list3 = []
        let m = list1.length
        let n = list2.length
        let i = 0
        let j = 0
        while(i < m && j < n) {
            if(list1[i] < list2[j]) {
                list3.push(list1[i])
                i++
            } else {
                list3.push(list2[j])
                j++
            }
        }
        while(i < m) {
            list3.push(list1[i])
            i++
        }
        
        while(j < n) {
            list3.push(list2[j])
            j++
        }
        return list3
    }
    
    
    store_in_order(node) {
        let traverse = function(node, list) {
            if(node == null) {
                return list
            }
            traverse(node.left, list)
            list.push(node.data)
            traverse(node.right, list)
            return list
        }
        
        let list = []
        return traverse(node, list)
        
    }
}
