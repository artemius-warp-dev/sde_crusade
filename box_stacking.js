//User function Template for javascript

/**
 * @param {number[]} height
 * @param {number[]} width
 * @param {number[]} length
 * @param {number} n
 * @returns {number}
*/




class Box {
    constructor() {
        this.h = 0
        this.w = 0
        this.l = 0
    }
}



class Solution {
    
    maxHeight(height, width, length, n)
    {
        //your code here
        let len = n * 3
           //STEP-1:Make an array of box (to store-> h,w,l)
    
        let arr = new Array(len)   //store  all 3 orientations
        let index = 0;
        for(let i=0;i<n;++i)
        {
            arr[index] = new Box()
            arr[index].h = height[i]
            arr[index].w = Math.max(width[i],length[i])
            arr[index].l = Math.min(width[i],length[i])
            index+=1
            
            arr[index] = new Box()
            arr[index].h = width[i]
            arr[index].w = Math.max(length[i],height[i])
            arr[index].l = Math.min(length[i],height[i])
            index+=1
           
            arr[index] = new Box()
            arr[index].h = length[i]
            arr[index].w = Math.max(height[i],width[i])
            arr[index].l = Math.min(height[i],width[i])
            index+=1
        }
   
    
    //STEP-2: Sort in DSC order of height
    arr.sort( (a, b) =>  (b.w*b.l) - (a.w*a.l))
    //console.log(arr)
    // //STEP-3: Find LIS
    let dp = new Array(len)
    
    
    for(let i=0; i<len; i++)
        dp[i] = arr[i].h
    
    let res = dp[0]
    for(let i = 1; i < len; i++){
        res = Math.max(res,dp[i]);
        for(let j = i-1; j >= 0; j--){
            if(arr[i].l < arr[j].l && arr[i].w < arr[j].w){
                dp[i] = Math.max(dp[i], arr[i].h+dp[j]);
                res = Math.max(res,dp[i]);
            }
        }
    }
    return res;
    }
     
}
