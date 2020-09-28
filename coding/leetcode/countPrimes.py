class Solution:
    def enumeratePrimes(self, n: int) -> list:
        '''
        :type n: int
        :rtype: list of int
        '''

        if n == 1: # O(1)
            return [] # O(1)

        # prime = [True]*n
        prime = dict() # O(1)

        for i in range(2, n): # O(n)
            prime[i] = True # O(1)

        p = 2 # O(1)

        while p**2 <= n:
            i = 2*p # O(1)

            for i in range(i, n, p): # O(n)
                prime[i] = False # O(1)

            for j in range(p + 1, n): # O(n)
                if prime[j] == True: # O(1)
                    p = j # O(1)
                    break # O(1)
            # print(f"p: {p}")

        prime = {k:v for k,v in prime.items() if v == True } # O(n)
        prime = list(prime.keys()) 
        return len(prime)