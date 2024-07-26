def singleNumber(nums):
    result = 0
    for num in nums:
        result ^= num
    return result

if __name__ == "__main__":
    # Ввод чисел от пользователя
    user_input = input("Введите числа (через пробел): ")
    # Преобразуем строку ввода в список целых чисел
    nums = list(map(int, user_input.split()))

    # Нахождение уникального числа
    unique_number = singleNumber(nums)

    # Вывод результата
    print(f"Единственное число: {unique_number}")
