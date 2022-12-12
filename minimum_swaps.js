//User function Template for javascript

/**
 * @param {number[]} nums
 * @returns {number}
*/
class Solution
{
    //Function to find the minimum number of swaps required to sort the array.
	minSwaps(nums)
	{
		// code here
		let n = nums.length
		let arrpos = []
		for(let i=0; i<n; i++ ) {
		    arrpos.push([nums[i], i])
		}
		
		arrpos.sort(function(a,b){
		    return a[0] - b[0]
		})
		
		//console.log(arrpos)
		
		let visited = new Array(n).fill(false)
		
		let ans = 0
		
		for(let i =0; i<n; i++) {
		    if(visited[i] || arrpos[i][i] == i)
		        continue
		        
		    let cycle_size = 0
		    let j = i
		    while(!visited[j]) {
		        visited[j] = true
		        
		        j = arrpos[j][1]
		       // console.log(j, arrpos[j])
		        cycle_size++
		       // console.log(cycle_size)
		    }
		    
		    if(cycle_size > 0) {
		        ans += (cycle_size - 1)
		    }
		}
		
		
    //0 vis[0] = true 3  [1 5 2 4 3]  vis[3] = true 0 c=2
    //1 vis[1] = true 2  [1 2 5 4 3]  vis[2] = true 4 [1,2,3,4 5] vis[4] = true 1 c=3
    //2 vis[2] = true   
    //3 vis[3]
    //4 vis[4]
    
    
    //1 5 2 4 3
    //1 2 5 4 3
    //1 2 3 4 5
		
		return ans
	}
}
