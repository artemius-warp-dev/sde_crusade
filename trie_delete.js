//User function Template for javascript

/*
class TrieNode{
    constructor(){ 
        this.children = new Array(26);
        this.children.fill(null);
        // isEndOfWord is True if node represent the end of the word 
        this.isEndOfWord = false
    }
}
*/


    // function isEmpty(root) {
    //     for(let i=0; i<26; i++) {
    //         if(root.children[i] != null)
    //             return false
    //     }
    //     return true
    // }

/**
 * @param {TrieNode} root (root of trie tree)
 * @param {string} key (key to be inserted)
*/1
class Solution {
    

    
    deleteKey(root, key){
        //code here
        //console.log(root.children)
        //console.log(key)
        
        let remove = function(root, key, depth) {
            // if(root == null)
            //     return null
                
            if(depth == key.length) {
                //console.log(depth)            
                if(root.isEndOfWord)
                    root.isEndOfWord = false
 
                
                return root
            }
        
            
            let index = key[depth].charCodeAt(0) - 'a'.charCodeAt(0)
            //console.log(index)
            //console.log(root.children[index])
            root.children[index] = remove(root.children[index], key, depth + 1)
  
            return root
            
        }
       let r = remove(root, key, 0) 
       //console.log(r)
       return r
    }
    
    
       // insert(root, key) {
    //     let pCrawl = root
        
    //     for(let i=0; i<key.length; i++) {
    //         let index = key[i].charCodeAt(0) - 'a'.charCodeAt(0)
    //         if(pCrawl.children[index] == null) {
    //             pCrawl.children[index] = new TrieNode()
    //         }
    //         pCrawl = pCrawl.children[index]
    //     }
        
    //     pCrawl.isEndOfWord = true
        
    // }
    
    // search(root, key) {
    //     let pCrawl = root
    //     for(let i=0; i<key.length; i++) {
    //         let index = key[i].charCodeAt(0) - 'a'.charCodeAt(0)
    //         if(pCrawl.chidren[index] == null)
    //             return false
            
    //         pCrawl = pCrawl.children[index]
    //     }
    //     return (pCrawl != null && pCrawl.isEndOfWord)
    // }
    
 
    
}
