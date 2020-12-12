Day 1
=====

## Part 1

This is a classic 2 sum problem. On unsorted input this can be solved in linear
time with linear memory using hash map to keep track of seen numbers in the
input.

* https://leetcode.com/problems/two-sum/

## Part 2

Another classic 3 sum problem. This can be solved with quadratic time if you fix
one number and solve 2 sum problem. But it will consume quadratic memory. We can
solve that by keeping a hash map of already seen numbers in outer loop.

* https://leetcode.com/problems/3sum/

