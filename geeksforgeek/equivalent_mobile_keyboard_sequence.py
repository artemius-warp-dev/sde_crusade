class Solution:

    def printSequence(self,S):
        # code here
        numeric_alpabet = ["2", "22", "222",
       "3", "33", "333",
       "4", "44", "444",
       "5", "55", "555",
       "6", "66", "666",
       "7", "77", "777", "7777",
       "8", "88", "888",
       "9", "99", "999", "9999"]

        output = ""
        n = len(S)
        for i in range(n):
            if S[i] == ' ':
                output = output + "0"
            else:
                position = ord(S[i]) - ord('A')
                output = output + numeric_alpabet[position]
        return output


if __name__ == '__main__':
    string = "GFG"
    s = Solution()
    print(s.printSequence(string))
        


