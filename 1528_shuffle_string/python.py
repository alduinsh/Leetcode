def shuffle_string(s, indices):
    # Создаем список, чтобы хранить перемещенные символы
    shuffled = [''] * len(s)

    # Проходим по символам строки и перемещаем их в нужные позиции
    for i in range(len(s)):
        shuffled[indices[i]] = s[i]

    # Преобразуем список обратно в строку и возвращаем
    return ''.join(shuffled)

# Примеры использования
# Пример 1
s1 = "codeleet"
indices1 = [4, 5, 6, 7, 0, 2, 1, 3]
result1 = shuffle_string(s1, indices1)
print(result1)  # Вывод: "leetcode"