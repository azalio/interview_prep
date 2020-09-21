#!/usr/bin/env python3


class Solution:
    def rotate(self, matrix): # O(n^2)
        '''
        :type matrix: list of list of int
        :rtype: list of list of int
        '''

        """
        two steps
        1. a[x,y] = a[y,x]
        2. a[1,-1] = a[-1,1]

        1,2,3
        4,5,6
        7,8,9

        0,0 -> 0,0
        0,1 -> 1,0
        0,2 -> 2,0
        1,2 -> 2,1
        2,2 -> 2,2

        1,4,7
        2,5,8
        3,6,9
        """
        matrix_len = len(matrix) # O(1)

        def swap_diagonal(matrix: list, step: int) -> list: # O(n)
            """
            Идем по верху. В каждом цикле уменьшаем матрицу, отбрасывая верхний и левый край.
            То есть:
            1, 2, 3
               5, 6
                  9
            """
            if matrix_len <= 1: # O(1)
                return matrix # O(1)

            for i in range(matrix_len - step): # O(n)
                matrix[step][step + i], matrix[step + i][step] = matrix[step + i][step], matrix[step][step + i] #O(1)
            return matrix # O(1)

        def swap_horizontal(matrix: list) -> list: 
            """
            Тут просто меняем местами элементы в списке так, так чтобы 0 -> -1, 1 -> -2, 2 -> -3
            """
            for i in range(matrix_len): # O(n^2)
                for j in range(matrix_len//2): # O(n)
                    matrix[i][j], matrix[i][-(j+1)] = matrix[i][-(j+1)], matrix[i][j] # O(1)
            return matrix

        for step in range(matrix_len): # O(n)*O(n) -> O(n^2)
            matrix = swap_diagonal(matrix, step) # O(n)

        matrix = swap_horizontal(matrix) # O(n^2)

        return matrix

obj = Solution()


matrix = [[1,2,3], [4,5,6], [7,8,9]]
res = obj.rotate(matrix)
assert res == [[7,4,1], [8,5,2], [9,6,3]]

matrix = [
  [ 1,  2,  3, 4],
  [ 5,  6,  7, 8],
  [ 9, 10, 11, 12],
  [13, 14, 15, 16]
]

ans = [
 [13,  9, 5, 1],
 [14, 10, 6, 2],
 [15, 11, 7, 3],
 [16, 12, 8, 4]
]

res = obj.rotate(matrix)
assert res == ans


matrix = [
  [10, 20],
  [30, 40]
]

ans = [
 [30, 10],
 [40, 20]
]

res = obj.rotate(matrix)
assert res == ans