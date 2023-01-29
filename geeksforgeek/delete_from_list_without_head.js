//User function Template for javascript

/**
 * @param {Node} del
*/

/*
class Node{
    constructor(data){
        this.data = data;
        this.next = null;
    }
}

let head;

*/

//Function to delete a node without any reference to head pointer.
class Solution {
    
    deleteNode(del)
    {
        //your code here
        let tmp = del.next
        del.data = tmp.data
        del.next = tmp.next
       
        //console.log(del)
        
        
        
    }
}
