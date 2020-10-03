#!/usr/bin/env python3

class Solution:
    def validSudoku(self, board: list) -> bool:
        seen = dict() # space O(81) -> O(1)
        for row in range(9): # O(81) -> O(1)
            for col in range(9):
                # print(f"row: {row}, col: {col}")
                value = board[row][col]
                if value == 0:
                    continue
                row_seen = f"{row}({value})"
                col_seen = f"({value}){col}"
                inner_seen = f"{row}({value}){col}"
                if row_seen in seen or col_seen in seen or inner_seen in seen:
                    return False
                seen[row_seen] = True
                seen[col_seen] = True
                seen[inner_seen] = True
        return True



obj = Solution()
input =  [[5,3,0,0,7,0,0,0,0],
          [6,0,0,1,9,5,0,0,0],
          [0,9,8,0,0,0,0,6,0],
          [8,0,0,0,6,0,0,0,3],
          [4,0,0,8,0,3,0,0,1],
          [7,0,0,0,2,0,0,0,6],
          [0,6,0,0,0,0,2,8,0],
          [0,0,0,4,1,9,0,0,5],
          [0,0,0,0,8,0,0,7,9]]

res = obj.validSudoku(input)
assert res == True