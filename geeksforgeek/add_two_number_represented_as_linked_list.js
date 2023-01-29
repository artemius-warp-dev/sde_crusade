//User function Template for javascript

/**
 * @param {Node} first
 * @param {Node} second
 * @returns {Node}
*/

/*
class Node{
    constructor(data){
        this.data = data;
        this.next = null;
    }
}
*/

function reverseList(head) {
        //your code here    
        let next = null
        let cur = head
        let prev = null
        while(cur != null) {
            next = cur.next
            cur.next = prev
            prev = cur
            cur = next
            
        }
        head = prev
        return head
}

class Solution {
    //Function to add two numbers represented by linked list.
    addTwoLists(first, second)
    {
        //your code here
        first = reverseList(first)
        second = reverseList(second)
        //console.log(first)
        let carry = 0
        
        let dummy = new Node(0)
        let temp = dummy
        while(first != null || second != null || carry == 1) {
            let sum = 0
            if(first != null) {
                sum = sum + first.data
                first = first.next
            }
            if(second != null) {
                sum = sum + second.data
                second = second.next
            }
            
            
            sum = sum + carry
            carry = Math.floor(sum / 10)
            temp.next = new Node(sum % 10)
            temp = temp.next
        }
        return reverseList(dummy.next)
    }
}
