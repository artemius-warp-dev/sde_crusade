//User function Template for javascript

/**
 * @param {MinHeapNode} root
 * @param {string} s
 * @return {string} 
*/

// BUG in 81 test case on GFG

class Index {
    constructor() {
        this.i = 0
    }
}

class Solution {
  decode_file(root,s){
    // code here
    if(s.length == 1) return ""
    
    let traverse = function(root, s, index) {
        if(root.left == null && root.right == null) {
            return root.data
        }
    
        if(s[index.i] == 0) {
            index.i = index.i + 1
            return traverse(root.left, s, index)
        } else {
            index.i = index.i + 1
            return traverse(root.right, s, index)
        }
    }
    
    let ans = ""
    let index = new Index()
    
    //console.log(i)
    while(index.i < s.length) {
        ans += traverse(root, s, index)
        //console.log(i)
        //break
    }
    
    return ans
    
  }
}


//C++
/*Complete the function below 
Which contains 2 arguments 
1) root of the tree formed while encoding
2) Encoded String*/
//  char travel(MinHeapNode* root, string &s, int &i){
//     if(root->left == NULL && root->right == NULL){
//         return root->data;
//     }
//     if(s[i] == '0'){
//         i++;
//         return travel(root->left, s, i);
//     }
//     else{
//         i++;
//         return travel(root->right,s, i);
//     }
// }

// string decode_file(struct MinHeapNode* root, string s)
// {
//     string ans = "";
//     int i = 0;
//     while(i < s.length()){
//         ans += travel(root, s, i);
//     }
//     return ans;
// }
