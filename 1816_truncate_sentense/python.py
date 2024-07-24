def truncate_sentence(s: str, k: int) -> str:
    # Разбиваем строку на слова
    words = s.split()
    
    # Соединяем первые k слов обратно в строку
    return ' '.join(words[:k])

def main():
    # Примеры использования
    print(truncate_sentence("Hello how are you Contestant", 4))  # "Hello how are you"
    print(truncate_sentence("What is the solution to this problem", 4))  # "What is the solution"
    print(truncate_sentence("chopper is not a tanuki", 5))  # "chopper is not a tanuki"

if __name__ == "__main__":
    main()
