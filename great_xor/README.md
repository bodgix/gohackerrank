# The Great XOR

This project contains solution to the [The Great XOR](https://www.hackerrank.com/contests/w28/challenges/the-great-xor) challenge
on Hacker Rank

# About the solution

The key to the algorithm for this solution was realizing, that XOR will only result
in a larger number if the right hand side operand of XOR has a `1` in place where
the left-hand operand has a `0`.

In other words, only changing `0`'s to `1`'s increases a binary number.
