import functools

def myCompare(y ,x):
    print([x,y])
    x_, y_, tmpx, tmpy = x,y, x,y
    #? xy or yx
    # 54546 or 54654
    
    count_num_of_x, count_num_of_y = 0,0

    while tmpx > 0:
        tmpx//=10
        count_num_of_x+=1
    while tmpy > 0:
        tmpy//=10
        count_num_of_y+=1

        
    while count_num_of_x > 0:
        y_*=10
        count_num_of_x-=1

    while count_num_of_y > 0:
        x_*=10
        count_num_of_y-=1

    # xyyx yxxy
    x_ += y 
    y_ += x
    
    print([x_, y_])
    #print(x_ < y_)
    if x_ < y_:
        return 1
    elif x_ > y_:
        return -1
    else:
        return 0

class Solution:

    


    

    def printLargest(self,arr):
        size = len(arr)
        arr = [int(elem) for elem in arr]
        arr.sort(key=functools.cmp_to_key(myCompare))
        arr.reverse()
        return ''.join(map(str, arr))

if __name__ == '__main__':
    a = ['54', '60', '546', '548']
    s = Solution()
    print(s.printLargest(a))

   # RIGTH SOLUTION FOR TIME TESTS ON GEEKSFORGEEK
   #  class Solution:
   # def printLargest(self,arr):
   #     def fun(a,b):
   #         x=a+b
   #         y=b+a
   #         if x>y:
   #             return 1
   #         else:
   #             return -1
   #     arr=sorted(arr,key=functools.cmp_to_key(fun),reverse=True)
   #     return "".join(arr)
