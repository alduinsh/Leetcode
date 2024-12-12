class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        known = set()    #set хранит уникальные объекты и позволяет решить O(n)
        for num in nums:
            if num in known:
                return True
            known.add(num)
        return False