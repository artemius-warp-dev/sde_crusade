class Solution:
    #Function to check if a string can be obtained by rotating
    #another string by exactly 2 places.
    def isRotated(self,str1,str2):
        #code here
        length = len(str1)
        if len(str2) != length or length < 2:
            return False
        anti_clock_rot = ""
        clock_rot = ""
        
        clock_rot = clock_rot + str1[2:] + str1[0:2]
        anti_clock_rot = anti_clock_rot + str1[length-2:] + str1[0:length-2]
        #print(clock_rot)
        #print(anti_clock_rot)
        return (anti_clock_rot == str2) or (clock_rot == str2)
        

    
if __name__ == '__main__':
    str1 = "amazon"
    str2 = "azonam"
    s = Solution()
    print(s.isRotated(str1, str2))
        
