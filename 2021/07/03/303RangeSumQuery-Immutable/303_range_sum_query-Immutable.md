# [303. Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/)


## 题目

Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

**Example:**

    Given nums = [-2, 0, 3, -5, 2, -1]
    
    sumRange(0, 2) -> 1
    sumRange(2, 5) -> -1
    sumRange(0, 5) -> -3

**Note:**

1. You may assume that the array does not change.
2. There are many calls to sumRange function.


## 题目大意

给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

示例：

```
给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3

```