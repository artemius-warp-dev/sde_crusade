//User function Template for javascript

/**
 * @param {number[]} nums
 * @return {number[]} 
*/
class Solution {
  	sortedArrayToBST(nums){
  		//code here
  		let res = []
  		if(nums.length == 0) return res
  		
  		let build_bst = function(nums, res, left_bound, right_bound) {
  		    
  		    if(left_bound > right_bound)
  		        return []
  		    
  		    let midpoint = Math.floor(left_bound + (right_bound - left_bound) / 2)
  		    //console.log(midpoint)
  		    res.push(nums[midpoint])
  		    build_bst(nums, res, left_bound, midpoint - 1)
  		    build_bst(nums, res, midpoint + 1, right_bound)
  		    return res
  		}
  		return build_bst(nums, res, 0, nums.length - 1)
  		
  	}
}
