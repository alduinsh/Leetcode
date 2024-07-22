'''_summary_
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false'''

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        return sorted(s) == sorted(t)

# Создаем объект класса Solution
solution = Solution()

# Проверяем работоспособность метода isAnagram
print(solution.isAnagram("listen", "silent")) 
print(solution.isAnagram("hello", "bello")) 