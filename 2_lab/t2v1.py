import numpy as np

# Параметры кода (23, 18) - длина кодовой комбинации и количество информационных разрядов
n = 23
k = 18

# Образующая матрица для (23, 18) кода
info_matrix = np.eye(k, dtype=int)  # единичная матрица размером 18x18 для информационных разрядов
parity_matrix = np.random.randint(0, 2, (k, n - k))  # случайная проверочная матрица 18x5 
generator_matrix = np.hstack((info_matrix, parity_matrix))  # полная образующая матрица 18x23

# Функция для кодирования

"""
Для кодирования используется умножение 
информационного вектора на матрицу GG с последующим взятием по модулю 2
"""
def encode(info_part, generator_matrix):
    codeword = np.mod(info_part @ generator_matrix, 2)
    return codeword

# Функция для проверки синдрома
"""_
Синдром ошибки вычисляется умножением кодового слова на транспонированную проверочную матрицу:
Если синдром = 0, то кодовое слово корректно
"""
def check_codeword(codeword, parity_check_matrix):
    syndrome = np.mod(codeword @ parity_check_matrix.T, 2)  # Убедимся, что используем транспонированную матрицу
    return syndrome

# Информационная часть (18 бит)
info_part = np.random.randint(0, 2, k)
print("Информационная часть:", info_part)

# Кодирование
codeword = encode(info_part, generator_matrix)
print("Сформированный кодовый вектор:", codeword)

# Формирование корректной проверочной матрицы
parity_check_matrix = np.hstack((parity_matrix.T, np.eye(n - k, dtype=int)))  # 5x23

# Проверка синдрома для выявления ошибок
syndrome = check_codeword(codeword, parity_check_matrix)
print("Синдром:", syndrome)

# Проверка наличия ошибок
if not np.any(syndrome):
    print("Ошибок нет.")
else:
    print("Обнаружена ошибка, синдром:", syndrome)
