//User function Template for javascript

/**
 * @param {Node} node
 * @return {Node} node
*/


/*
  Pairwise swap a linked list
  The input list will have at least one element
  node is defined as

class Node{
    constructor(data){
        this.data = data;
        this.next = null;
    }
}

let head;
This is method only submission.
You only need to complete the below method.
*/

class Solution {
    pairWiseSwap(node)
    {
        //your code here
        let p_1 = node
        let ns = p_1
        if(p_1.next == null)
            return ns
        else
            ns = p_1.next
        let p_2 = p_1
        let tmp = null
        while(true) {
        p_2 = p_1.next
        tmp = p_2.next
        p_2.next = p_1
        if(tmp == null || tmp.next == null) {
            p_1.next = tmp
            break
        }
        p_1.next = tmp.next
        p_1 = tmp
        }
        return ns
    }
}
