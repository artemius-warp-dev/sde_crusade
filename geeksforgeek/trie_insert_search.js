//User function Template for javascript



/**
 * 
 * insert 
 * @param {TrieNode} root (root of trie tree)
 * @param {string} key (key to be inserted)
 * 
 * search
 * @param {TrieNode} root (root of trie tree)
 * @param {string} key (key to be searched)
 * @returns {boolean}  (true if key presents in trie, else false)
*/
class Solution 
{
    //Function to insert string into TRIE.
    insert(root, key) 
    {
       // Your code here
       let node = root
       for(let level = 0; level<key.length; level++) {
           let index = key[level].charCodeAt(0) - 'a'.charCodeAt(0);
           if(node.children[index] == null) {
               node.children[index] = new TrieNode()
           }
           node = node.children[index]
       }
       node.isEndOfWord = true
    }
    
    //Function to use TRIE data structure and search the given string.
    search(root, key) 
    {
        // Your code here
        let node = root
        for(let level = 0; level<key.length; level++) {
           let index = key[level].charCodeAt(0) - 'a'.charCodeAt(0);
           if(node.children[index] == null) {
               return 0
           }
           node = node.children[index]
        }
        return node.isEndOfWord
    }
}
