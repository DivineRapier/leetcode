# 992. [Subarrays with K Different Integers](https://leetcode.com/problems/subarrays-with-k-different-integers/)

Given an array `A` of positive integers, call a (contiguous, not necessarily distinct) subarray of `A` good if the number of different integers in that subarray is exactly `K`.

(For example, `[1,2,3,1,2]` has `3` different integers: `1`, `2`, and `3`.)

Return the number of good subarrays of `A`.

 

**Example 1:**
```
Input: A = [1,2,1,2,3], K = 2
Output: 7
Explanation: Subarrays formed with exactly 2 different integers: [1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].
```

**Example 2:**
```
Input: A = [1,2,1,3,4], K = 3
Output: 3
Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].
```

**Note:**
1. `1 <= A.length <= 20000`
1. `1 <= A[i] <= A.length`
1. `1 <= K <= A.length`

## Tips

题目要求 `exactly K` 个不同元素组成的子数组，可以转换为 `at most K` - `at most K-1`。所以，问题转换为: 连续的不同元素组成的子数组的个数。

使用双指针 `fast` `slow`，`fast` 每移动一次增加 `fast` - `slow` + `1` (该数字表示当前位置数字与`slow`之间子数组的个数)

