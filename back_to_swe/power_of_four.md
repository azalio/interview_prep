# Power of four

Observe that the binary representation of a power of four contains exactly one bit turned on.  
`1 (represents 1 under base 2)`
`100 (represents 4 under base 2)`
`10000 (represents 16 under base 2)`
`1000000 (represents 64 under base 2)`

Moreover, this one bit must occur in an odd position from the right.  
`1 (1st bit from the right)`
`100 (3rd bit from the right)`
`10000 (5th bit from the right)`
`1000000 (7th bit from the right)`

Thus, it suffices to check whether the provided integer has exactly one bit turned on in an odd position.  
Letting x denote our input, we can first verify that our integer has only one bit turned on by verifying that `x` is nonzero and `x & (x - 1)` equals zero.  
Example of `x & (x - 1)`:
`x = 10000` (represents 16 under base 2)
`x - 1 = 1111` (represents 15 under base 2)
`10000&01111 = 00000` (test passes - the leading bit of x is set and the rest are 0)  

The reason why this works is explained in the solution of the "Power of Two" problem.

At this point, we know that our inputted integer has only one bit on.  
Now we just need to verify that this bit that is toggled on takes place in an odd position.
We can do this by verifying that the logical AND of x and `0x55555555` evaluates to true.
Why? Because the binary representation of the hexadecimal number `55555555` is `1010101010101010101010101010101`.
In other words, this hexadecimal number represents the binary number whose odd positions are toggled on and whose even positions are toggled off.
Thus, the boolean expression (x & 0x55555555) will evaluate to true if and only if x's only on bit is in an odd position.

## То есть

- Проверили что больше нуля
- Убедились что это степень двойки через битовое `AND` на входящем числе и числа перед ним
(это работает потому что это двоичная система счисления и степень двойки всегда получается с ведущим битом)

- убедились что бит выставлен в четной позиции (именно так определяется степень четверки в двоичной системе)
