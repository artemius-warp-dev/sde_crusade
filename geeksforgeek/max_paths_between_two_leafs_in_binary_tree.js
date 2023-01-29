//User function Template for javascript

/*
class Node
{
    constructor(x){
        this.key=x;
        this.left=null;
        this.right=null;
    }
}
*/



//   int helper(Node *root,int &sum)
//     {
//         if(!root) return INT_MIN;
//         if(root && !root->left && !root->right) return root->data;
//         int l=helper(root->left,sum);
//         int r=helper(root->right,sum);
//         if(root->left && root->right)
//         {
//             sum=max(sum,l+r+root->data); 
//         }
//         return root->data+max(l,r);
//     }
    
//     int maxPathSum(Node* root)
//     {
//       int sum=INT_MIN;
//       int a=helper(root,sum);
//       if(!root->left || !root->right)
//       return max(a,sum);
//       return sum;
//     }

/**
 * @param {Node} root
 * @return {number}
*/

class Res {
    constructor() {
        this.value = Number.NEGATIVE_INFINITY
    }
}

class Solution {
  	maxPathSum(root){
  		//code here
  		
  		let max_path_sum = function(node, maxi) {
  		   
  		    
  		    if(node == null) {
  		        return Number.NEGATIVE_INFINITY
  		    }
  		    
  		    if(node != null && node.left == null && node.right == null) {
  		        return node.key
  		    }
  		    
  		    let left = max_path_sum(node.left, maxi)
  		    let right = max_path_sum(node.right, maxi)
  		    //console.log(left,right, maxi)
  		    //console.log(node.left, node.right)
  		    if(node.left != null && node.right != null) {
  		        //console.log([maxi.value, left + right + node.key])
  		        maxi.value = Math.max(maxi.value, left + right + node.key)
  		        //console.log(maxi.value)
  		       
  		    }
  		    //console.log(left, right, node.key)
  		    return Math.max(left, right) + node.key
  		    
  		    
  		}
  		
  		let maxi = new Res()
  		let a = max_path_sum(root, maxi)
  		//console.log(a, maxi.value)
  		if(root.left == null || root.right == null) 
  		    return Math.max(a,maxi.value)
  		
  		return maxi.value
  		
  	}
}
