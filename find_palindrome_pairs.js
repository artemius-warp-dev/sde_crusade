//User function Template for javascript

const ALPHABET_SIZE = 26

function is_palindrome(str, i, len) {
     while(i < len) {
        if(str[i] != str[len]) 
            return false
        i++, len--
     }
     return true
 }   

 class Node {
     constructor() {
        this.children = new Array(ALPHABET_SIZE).fill(null)
        this.isLeaf = false
        this.pos = []
        this.index = -1
     }
     
     
     

 }
 
 class Trie {
     constructor() {
         this.root = new Node()
     }
     
     insert(key, word_index) {
         let p_crawl = this.root
         
         
         for(let level=key.length-1; level>=0; level--) {
            let ch_index = key[level].charCodeAt(0) - 'a'.charCodeAt(0)
            if(!p_crawl.children[ch_index]) {
                p_crawl.children[ch_index] = new Node()
            }
            
            if(is_palindrome(key, 0, level)) {
                p_crawl.pos.push(word_index)
            }
            
            p_crawl = p_crawl.children[ch_index]
         }
         
         p_crawl.index = word_index
         p_crawl.pos.push(word_index)
         p_crawl.isLeaf = true
     }
     
     
     search(key, word_index, result) {
         let p_crawl = this.root
         for(let level = 0; level < key.length; level++) {
             let ch_index = key[level].charCodeAt(0) - 'a'.charCodeAt(0)
             if(p_crawl.index >= 0 && p_crawl.index != word_index && is_palindrome(key, level, key.length - 1)) {
                 result.push([word_index, p_crawl.index])
             }
             
            if(!p_crawl.children[ch_index]) {
                //console.log(p_crawl.children[ch_index])
                return
            }
            
            p_crawl = p_crawl.children[ch_index]
         }
         
         
         for(let i of p_crawl.pos) {
            if(i == word_index)
                continue
            result.push([i, word_index])
         }
     }
     
     
 }   


/**
 * @param {number} N
 * @param {number[]} arr
 * @return {boolean} 
*/

class Solution {
  palindromepair(N,arr){
    //code here
    
    if(arr.length == 1) return false
    
    let trie = new Trie()
    for(let i=0; i<N; i++) {
        trie.insert(arr[i], i)
    }
    
    let result = new Array()
    for(let i = 0; i<arr.length; i++) {
        trie.search(arr[i], i, result)
        if(result.length > 0)  {
            //console.log(result)
            return true
        }
    }
    return false
  }
}
