import numpy as np

# Параметры кода (23, 18)
n = 23  # длина кодового слова
k = 18  # длина информационной части

# Создание образующей матрицы
info_matrix = np.eye(k, dtype=int)  # Единичная матрица 18x18
parity_matrix = np.random.randint(0, 2, (k, n - k))  # Случайная проверочная матрица 18x5
generator_matrix = np.hstack((info_matrix, parity_matrix))  # Полная образующая матрица 18x23

# Формирование проверочной матрицы
parity_check_matrix = np.hstack((parity_matrix.T, np.eye(n - k, dtype=int)))  # Проверочная матрица 5x23

# Кодирование
def encode(info_part, generator_matrix):
    """
    Кодирование выполняется умножением информационного вектора
    на генераторную матрицу с последующим взятием по модулю 2.
    """
    codeword = np.mod(info_part @ generator_matrix, 2)
    return codeword

# Проверка синдрома
def check_codeword(codeword, parity_check_matrix):
    """
    Синдром ошибки вычисляется умножением кодового слова
    на транспонированную проверочную матрицу.
    """
    syndrome = np.mod(codeword @ parity_check_matrix.T, 2)
    return syndrome

# Исправление ошибок
def correct_errors(codeword, syndrome, parity_check_matrix):
    """
    Попытка исправить одиночные и двойные ошибки:
    - Одиночная ошибка: синдром соответствует столбцу проверочной матрицы.
    - Двойная ошибка: синдром равен сумме двух столбцов проверочной матрицы.
    """
    # Проверка одиночной ошибки
    for i, col in enumerate(parity_check_matrix.T):
        if np.array_equal(syndrome, col):
            print(f"Обнаружена одиночная ошибка в разряде {i}. Исправляем.")
            codeword[i] = 1 - codeword[i]  # Инверсия бита
            return codeword

    # Проверка двойной ошибки
    for i, col_i in enumerate(parity_check_matrix.T):
        for j, col_j in enumerate(parity_check_matrix.T):
            if i < j and np.array_equal(syndrome, np.mod(col_i + col_j, 2)):
                print(f"Обнаружена двойная ошибка в разрядах {i} и {j}. Исправляем.")
                codeword[i] = 1 - codeword[i]
                codeword[j] = 1 - codeword[j]
                return codeword

    print("Ошибка не обнаружена или она превышает допустимые пределы (2 ошибки).")
    return codeword

# Генерация информационной части
info_part = np.random.randint(0, 2, k)
print("Информационная часть:", info_part)

# Кодирование
codeword = encode(info_part, generator_matrix)
print("Сформированный кодовый вектор:", codeword)

# Введение ошибки (одиночной или двойной)
codeword_with_error = codeword.copy()
codeword_with_error[2] = 1 - codeword_with_error[2]  # Одиночная ошибка
codeword_with_error[7] = 1 - codeword_with_error[7]  # Вторая ошибка
print("Кодовое слово с ошибками:", codeword_with_error)

# Проверка синдрома
syndrome = check_codeword(codeword_with_error, parity_check_matrix)
print("Синдром:", syndrome)

# Исправление ошибок
corrected_codeword = correct_errors(codeword_with_error, syndrome, parity_check_matrix)
print("Исправленное кодовое слово:", corrected_codeword)

# Проверка на корректность после исправления
final_syndrome = check_codeword(corrected_codeword, parity_check_matrix)
if not np.any(final_syndrome):
    print("Ошибки исправлены. Кодовое слово корректно.")
else:
    print("Ошибки не удалось исправить.")
