//User function Template for javascript

/**
 * @param {number} N
 * @param {number} C
 * @param {number[]} pages
 * @return {number}
 */

class Solution {
  pageFaults(N, C, pages) {
    //code here
    let memory = new Set()
    let cache = new Map()
    let page_fault = 0
    for(let i=0; i<N; i++) {
        if(memory.size < C) {
            if(!memory.has(pages[i])) {
                memory.add(pages[i])
                page_fault++
            }
            cache[pages[i]] = i
        } else {
            //LRU algo
            if(!memory.has(pages[i])) {
                page_fault++
                let lru = Number.MAX_VALUE
                let value
                for(let it of memory) {
                    if(cache[it] < lru) {
                        lru = cache[it]
                        value = it
                    }
                }
                
               
                //console.log(value)
                memory.delete(value)
                //console.log(memory)
                memory.add(pages[i])
                cache[pages[i]] = i
            }
            cache[pages[i]] = i
        }
        
    }
    return page_fault
  }
}
