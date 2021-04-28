# [290. Word Pattern](https://leetcode.com/problems/word-pattern/)

## 题目

Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

Example 1:

```c
Input: pattern = "abba", str = "dog cat cat dog"
Output: true
```

Example 2:

```c
Input:pattern = "abba", str = "dog cat cat fish"
Output: false
```

Example 3:

```c
Input: pattern = "aaaa", str = "dog cat cat dog"
Output: false
```

Example 4:

```c
Input: pattern = "abba", str = "dog dog dog dog"
Output: false
```

Note:

You may assume pattern contains only lowercase letters, and str contains lowercase letters separated by a single space.
