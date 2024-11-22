import numpy as np

# Информационная часть образующей матрицы кода
info_matrix = np.array([
    [0, 0, 0, 0, 1],
    [0, 0, 0, 1, 0],
    [0, 0, 1, 0, 0],
    [0, 1, 0, 0, 0],
    [1, 0, 0, 0, 0]
])

print("Матрица кода")
print(info_matrix)
print("")
print("Проверочная матрица")

# Проверочная матрица
check_matrix = np.array([
    [1, 0, 1, 0],
    [0, 0, 1, 1],
    [1, 1, 1, 0],
    [1, 1, 0, 1],
    [0, 1, 1, 1]
])
print(check_matrix)
print(" ")
# Полная образующая матрица
parity_matrix_full = np.hstack((info_matrix, check_matrix))

# Информационная часть кода
info_part = np.array([1, 0, 1, 0, 1]) #1 0 1 0 1

# Формирование избыточного кода (вычисляем контрольные разряды) умножение векторов и взятие остатка по mod 2
redundant_code = np.mod(info_part @ check_matrix, 2)

# Полный кодовый вектор. объединение инф. частии контрольных разрядов
code_word = np.concatenate((info_part, redundant_code))
print("Сформированный кодовый вектор:", code_word)

# Проверка кодовой комбинации
"""Синдром ошибки вычисляется как произведение кодового слова
на транспонированную проверочную матрицу с последующим взятием 
остатка по модулю 2:"""

def check_codeword(codeword, check_matrix):
    # Вычисляем синдром
    syndrome = np.mod(codeword @ check_matrix, 2)
    print("Синдром:", syndrome)
    return syndrome

# Пример проверочных комбинаций (Комбинация 1 и Комбинация 2)
codeword_1 = np.array([0, 1, 0, 1, 0, 1, 1, 0, 0 ]) #0 1 0 1 0 1 1 0 0 
codeword_2 = np.array([1, 0, 0, 1, 0, 1, 1, 1, 0]) #1 0 0 1 0 1 1 1 0

# Проверка комбинаций
print("Проверка Комбинации 1:")
syndrome_1 = check_codeword(codeword_1, parity_matrix_full.T)

print("\nПроверка Комбинации 2:")
syndrome_2 = check_codeword(codeword_2, parity_matrix_full.T)

# Поиск ошибки для однократной ошибки
"""
Если синдром ошибки не равен нулю, то 
он указывает на столбец проверочной матрицы HTHT, 
соответствующий ошибочному разряду
"""
def find_error(syndrome):
    for i, row in enumerate(parity_matrix_full.T):
        if np.array_equal(syndrome, row):
            print(f"Ошибка в разряде {i}")
            return i
    print("Ошибка не обнаружена или она двукратная.")
    return -1

print("\nПоиск ошибки для Комбинации 1:")
find_error(syndrome_1)

print("\nПоиск ошибки для Комбинации 2:")
find_error(syndrome_2)
