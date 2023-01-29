class Solution:


    def romanToDecimal(self, S): 
        roman = {}
        roman['I'] = 1
        roman['V'] = 5
        roman['X'] = 10
        roman['L'] = 50
        roman['C'] = 100
        roman['D'] = 500
        roman['M'] = 1000
    
        # code here
        l = len(S)
        i = 0
        res = 0
        while i < l:
            #value = S[i]
            #next_value = S[i+1]
            if i != (l - 1) and  (roman[S[i]] < roman[S[i+1]]):
                res = res + (roman[S[i+1]] - roman[S[i]])
                i = i + 2
            else:
                res = res + roman[S[i]]
                i = i+1

        return res
        




if __name__ == '__main__':
    str1 = "III"
    s = Solution()
    print(s.romanToDecimal(str1))
        
