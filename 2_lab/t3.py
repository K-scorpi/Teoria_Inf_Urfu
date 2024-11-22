import numpy as np
from itertools import combinations

# Новый образующий многочлен (матрица)
generator_polynomial = np.array([1, 0, 0, 1, 1, 1])  # (1 + x^3+ x^5) 1 0 0 1 1 1
n = 14  # Общая длина кодовой комбинации (избыточный код)
k = 9  # Длина информационной части

# Информационная часть кода
information_bits = np.array([1, 0, 0, 1, 1, 1, 0, 1, 0]) #1 0 0 1 1 1 0 1 0

# Функция для умножения многочленов по модулю 2
def poly_mult(p1, p2):
    result = np.zeros(len(p1) + len(p2) - 1, dtype=int)
    for i in range(len(p1)):
        for j in range(len(p2)):
            result[i + j] ^= p1[i] & p2[j]
    return result


# Функция для деления многочленов по модулю 2
def poly_dividend(dividend, divisor):
    n = len(dividend)
    m = len(divisor)
    quotient = np.zeros(n - m + 1, dtype=int)
    remainder = np.copy(dividend)

    for i in range(n - m + 1):
        if remainder[i] == 1:
            quotient[i] = 1
            for j in range(m):
                remainder[i + j] ^= divisor[j]

    return quotient, remainder


# Преобразование информационной части в кодовое слово (избыточный код)
def encode_codeword(information_bits, generator_polynomial):
    # Добавляем нули в конце информационного сообщения
    information_with_zeros = np.hstack([information_bits, np.zeros(len(generator_polynomial) - 1, dtype=int)])

    # Делаем деление, чтобы найти остаток (избыточные биты)
    _, remainder = poly_dividend(information_with_zeros, generator_polynomial)

    # Получаем избыточный код
    codeword = np.hstack([information_bits, remainder])
    return codeword


# Проверка кодовой комбинации
def check_codeword(codeword, generator_polynomial):
    _, remainder = poly_dividend(codeword, generator_polynomial)
    return remainder


# Преобразуем информационную часть в кодовое слово
codeword = encode_codeword(information_bits, generator_polynomial)

print("Избыточный код:", codeword)

# Проверка кодовой комбинации
remainder = check_codeword(codeword, generator_polynomial)
print("Остаток при проверке кодовой комбинации:", remainder)


# Определение однократной ошибки
def check_single_error(codeword, generator_polynomial):
    for i in range(len(codeword)):
        codeword_with_error = np.copy(codeword)
        codeword_with_error[i] ^= 1  # Вставляем ошибку в i-ый разряд
        remainder = check_codeword(codeword_with_error, generator_polynomial)

        # Обрезаем остаток до длины n - k (в данном случае 5)
        remainder = remainder[:n - k]  # Убедимся, что остаток имеет правильную длину

        if np.all(remainder == np.zeros(n - k, dtype=int)):  # Исправление сравнения с размерностью остатка
            print(f"Однократная ошибка может быть в разряде {i + 1} (индексация с 1).")


# Пример однократных ошибок
check_single_error(codeword, generator_polynomial)


# Примеры двукратных ошибок
def check_double_error(codeword, generator_polynomial):
    for i, j in combinations(range(len(codeword)), 2):
        codeword_with_errors = np.copy(codeword)
        codeword_with_errors[i] ^= 1
        codeword_with_errors[j] ^= 1
        remainder = check_codeword(codeword_with_errors, generator_polynomial)

        # Обрезаем остаток до длины n - k (в данном случае 5)
        remainder = remainder[:n - k]  # Убедимся, что остаток имеет правильную длину

        if np.all(remainder == np.zeros(n - k, dtype=int)):  # Исправление сравнения с размерностью остатка
            print(f"Двукратная ошибка возможна в разрядах {i + 1} и {j + 1} (индексация с 1).")


# Пример двукратных ошибок
check_double_error(codeword, generator_polynomial)
