class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:  # Добавлен self
        sdict = {}
        tdict = {} 
        for i in range(len(s)):
            if s[i] in sdict:
                if sdict[s[i]] != t[i]:
                    return False     
            else: 
                if t[i] in tdict and tdict[t[i]] != s[i]:
                    return False
                else:
                    sdict[s[i]] = t[i]
                    tdict[t[i]] = s[i]
        return True
    
sol = Solution()
print(sol.isIsomorphic("egg", "add"))  # Вернет True