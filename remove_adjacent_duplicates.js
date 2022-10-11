//User function Template for javascript



/**
 * @param {string} s
 * @return {string}
 */

class Solution {
    rremove_util(s, last_removed) {
       
         
        
        
       // console.log(s.length)
       
        if(s.length  <= 1) return [s, ""]
        let res
        let rem_str
        
        
        if(s[0] == s[1]) {
            last_removed = s[0]
            while(s.length > 1 && s[0] == s[1]) {
                s = s.substr(1,)
            }
            res = s.substr(1,)
            //console.log(res.length)
            return this.rremove_util(res, last_removed)
        } else {
            //console.log(["input",s])
            [rem_str, last_removed] = this.rremove_util(s.substr(1,),last_removed)
            //console.log(rem_str)


            if(rem_str.length != 0 && rem_str[0] == s[0]) {
                last_removed = s[0]
                return [rem_str.substr(1,), last_removed]
                
            }
        }

        //console.log([s, rem_str, last_removed])

        if(s[0] == last_removed) {
            //console.log([s, rem_str, last_removed])
            return [rem_str, last_removed]
        }

        //console.log(s[0] + rem_str)
        return [s[0] + rem_str, ""]
        //console.log([rem_str, s, last_removed])
        //console.log(s)
        
        //let rem_str = this.rremove_util(s.substr(1,), last_removed) 
        //console.log(["output", res])
       

        //console.log(rem_str)
        //console.log(last_removed)
       
        // if(s[0] == last_removed) {
        //     return rem_str
        // } else {
            
        // }

        //console.log([rem_str, last_removed])
        //console.log([rem_str, s])
        
       
        //console.log([s.charCodeAt(0), last_removed])
       
        //console.log([s, rem_str])


         //res = s[0] + rem_str
        //console.log([res, last_removed])
        
        //return res
        //  }
        

        
        
    }

    rremove(s) {
      //code here
        let last_removed = ""
        let res
        [res, last_removed] = this.rremove_util(s, last_removed)
        return res
      
      
  }
}


let s = new Solution()
let str = "qzcznpppqdssfxnttifmaddomiwwjiqyveunauupttttfaiialuueuuuuuaxrroniggbb"
//s.rremove(str)
//console.log("END")
let res = s.rremove(str)
console.log(res)
console.log(res == "qzcznqdfxnifmaomijiqyveunapfleaxoni")


// class Solution {
//  rremove(s) {
     
//    function traverseString(oldStr){
//        if(oldStr === '') return '';
//        let newStr = '';
//        let temp = '';
//        let count = 1;
       
//        for(let i=0; i<oldStr.length; i++){
//            if(temp === oldStr[i]){
//                count++;
//            } else {
//                if(count === 1){
//                   newStr += temp;
//                }
//                temp = oldStr[i];
//                count = 1;
//            }
//        }
       
//        if(count === 1)
//            newStr += temp;
       
//        if(newStr === oldStr) return oldStr;
       
//        return traverseString(newStr);
       
//    }
 
//    return traverseString(s);
//  }
 
// }
