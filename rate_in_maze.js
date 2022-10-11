/**
 * @param {number[][]} m
 * @param {number} n
 * @return {string[]}
 */

class Solution {
   findPath(m,n){
       //code here

       // visited node matrix
       let rows = m.length ;
       let columns = m[0].length;
       let visited = new Array(rows);
       let result = [];
       for(let i=0;i<rows;i++){
           visited[i] = new Array(columns).fill(0);
       }

       if(m[0][0]==0) return [];
       helper(m,visited,"",0,0);
       
       return result.sort();


       function helper(m,visited ,s,row,col){
           // base case
           if(row==n-1 && col==n-1) {
               result.push(s);
               return;
           }
           else{
               // mark the cell as visited
               visited[row][col] = 1;
           }
           
           //down
           if(row+1<n && visited[row+1][col]==0 && m[row+1][col]==1){
               helper(m,visited,s+"D",row+1,col);
           }
           //left
           if(col-1>=0 && visited[row][col-1]==0 && m[row][col-1]==1){
               helper(m,visited,s+"L",row,col-1);
           }
           //right
           if(col+1<n && visited[row][col+1]==0 && m[row][col+1]==1){
               helper(m,visited,s+"R",row,col+1);
           }
           //up
           if(row-1>=0 && visited[row-1][col]==0 && m[row-1][col]==1){
               helper(m,visited,s+"U",row-1,col);
           }
           
           // mark this node as unvisited
           visited[row][col] = 0;
       }


       
   }
}

let s = new Solution()
let N = 4
let matrix = [[1,0,0,0], [1,1,0,1], [1,1,0,0], [0,1,1,1]]
console.log(s.findPath(matrix, N))
