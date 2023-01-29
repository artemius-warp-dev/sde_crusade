
class PetrolPump{
    constructor(p, d){
        this.petrol = p;
        this.distance = d;
    }
}





/**
 * @param {PetrolPump[]} p
 * @param {number} n
 * @returns {number}
*/

class Solution {
    //Function to find starting point where the truck can start to get through
    //the complete circle without exhausting its petrol in between.
    tour(p, n)
    {
        //your code here
        let balance = 0
        let start = 0
        let sumRemaining = 0
        let gas_index = 0
        let distance_index = 1
        
        for(let i = 0; i < n; i++) {
            let remaining = p[i].petrol - p[i].distance
            //console.log(remaining)
            if(sumRemaining >= 0)
                sumRemaining += remaining
            else {
                sumRemaining = remaining
                start = i
            }
            
            balance += remaining
            gas_index+=2
            distance_index+=2
            //console.log([balance, remaining])
        }

        console.log(balance)
        if(balance >= 0)
            return start
        else
            return -1
    }
}

let s = new Solution()
let n = 4
let p = [3, 6, 6, 5, 7, 3, 4, 5] 
let arr = new Array(n)
for(let i=0;i<2*n;i+=2)
     arr[i/2] = new PetrolPump(p[i], p[i+1]);



console.log(s.tour(arr,n))
