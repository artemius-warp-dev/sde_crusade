//User function Template for javascript

/**
 * @param {number} N
 * @param {number} M
 * @param {number[][]} A
 * @param {number} X
 * @param {number} Y
 * @return {number}
 */
 
 class Node {
    constructor(r,c, distance) {
        this.row = r
        this.col = c
        this.distance = distance
    }
}

 

class Solution {
  shortestDistance(N, M, A, X, Y) {
    //code here

    let pq = []
    let shortest_distance_matrix = new Array(N).fill(Number.MAX_VALUE).map(() => new Array(M).fill(Number.MAX_VALUE))
    if(A[0][0] == 0) return -1
    shortest_distance_matrix[0][0] = 0
    
    pq.push(new Node(0,0,0))
    
    let dr = [-1, 0, 1, 0]
    let dc = [0, 1, 0, -1]
        
    while(pq.length != 0) {
        //console.log(pq.harr[0])
        let node = pq.shift()
        
        let dis = node.distance
        let r = node.row
        let c = node.col
        
        for(let i =0; i<4; i++) {
            let nrow = r + dr[i]
            let ncol = c + dc[i]
            if(nrow >=0 && ncol >= 0 && nrow < N && ncol < M && dis + 1 < shortest_distance_matrix[nrow][ncol] && A[nrow][ncol] == 1) {
                shortest_distance_matrix[nrow][ncol] = dis + 1
                //console.log([nrow, X], [ncol, Y])
                if(nrow == X && ncol == Y)
                    return dis + 1
                pq.push(new Node(nrow, ncol, dis+1))
            }
        }
        
       
    }
    return -1

  }
}
