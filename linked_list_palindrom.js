//User function Template for javascript

/**
 * @param {Node} head
 * @returns {boolean}
*/

/*
class Node{
    constructor(data){
        this.data = data;
        this.next = null;
    }
}
*/

function reverse_list(head) {
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
    //Function to check whether the list is palindrome.
    isPalindrome(head)
    {
        //your code here
        
        if(head == null || head.next == null)
            return true
        
        let middle = null
        let slow = head
        let fast = head
        while(fast.next != null && fast.next.next != null) {
                slow = slow.next
                fast = fast.next.next
        }
        
        middle = slow
        middle.next = reverse_list(middle.next)
        let right_ptr = middle.next
        
        while(right_ptr != null) {
            //console.log([head.data, right_ptr.data])
            if(head.data != right_ptr.data)
                return false
            head = head.next
            right_ptr = right_ptr.next
        }
        return true
    }
}
