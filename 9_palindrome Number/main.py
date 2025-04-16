class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False  # Отрицательные числа не могут быть палиндромами
        reverse = 0
        xcopy = x
        
        while x > 0:
            reverse = (reverse * 10) + (x % 10)
            x //= 10
        return reverse == xcopy

if __name__ == "__main__":
    solution = Solution()
    user_input = input("Введите числа (через пробел): ")
    numbers = user_input.split()
    
    for num in numbers:
        try:
            num_int = int(num)
            if solution.isPalindrome(num_int):
                print(f"{num_int} является палиндромом")
            else:
                print(f"{num_int} не является палиндромом")
        except ValueError:
            print(f"{num} не является целым числом")
            
            
class Solution:
    def isPalindrom(self, x: int) -> bool:
        if x < 0:
            return False
        reverse = 0
        xcopy = x
        
        while x > 0:
            reverse = (reverse * 10) + (x % 10)
            xcopy //= 10
        return reverse == xcopy
    
if __name__ == "main":
    solution = Solution()
    user_input = input('Введите числа без пробела')
    numbers = user_input.split()
    
    for num in numbers:
        try:
            num_int = int(num)
            if solution.isPalindrom(num_int):
                print(f"{num_int} является палиндромом")
            else:
                print(f"{num} не является палиндромом")
        except ValueError:
            print(f"{num} не является целым числом")