// User function Template for javascript

/**
 * @param {Node} root
 * @param {number} input
 * @returns {number}
*/

class Solution {
    // Function to return the ceil of given number in BST.
    findCeil(root, input) {
        // your code here
        let ceil = -1
        if(root.data == input) return root.data
        while(root != null) {
            if(root.data < input) {
                root = root.right
            } else {
                ceil = root.data
                root = root.left
            }
        }
        return ceil
    }
}
