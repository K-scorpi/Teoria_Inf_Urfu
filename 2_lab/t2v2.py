import numpy as np
from itertools import combinations

# Параметры кода
n = 15  # Длина кодового слова
t = 3   # Максимум исправляемых ошибок (трёхкратные ошибки)

# Образующая и проверочная матрица для примера
k = n - t - 1  # Размер информационной части
info_matrix = np.eye(k, dtype=int) #  единичная
parity_matrix = np.random.randint(0, 2, (k, n - k)) # случайная проверочная матрица 18x5 
generator_matrix = np.hstack((info_matrix, parity_matrix)) # полная образующая матрица 

# Проверочная матрица
parity_check_matrix = np.hstack((parity_matrix.T, np.eye(n - k, dtype=int)))

# Функция для генерации синдрома
"""
Синдром вычисляется как произведение 
вектора ошибок на транспонированную проверочную матрицу
"""
def generate_syndrome(error_vector, parity_check_matrix):
    return np.mod(error_vector @ parity_check_matrix.T, 2)

# Составление таблицы опознавателей
"""
    Перечисление всех комбинаций t=3 позиций из n=15.
    Например, для позиций ошибок 1,3,5 создаётся вектор:
    (0,1,0,1,0,1,0,0,0,0,0,0,0,0,0)
"""
error_patterns = list(combinations(range(n), t))  # Все трёхкратные ошибки
syndrome_table = {}

for error_positions in error_patterns:
    error_vector = np.zeros(n, dtype=int)
    for pos in error_positions:
        error_vector[pos] = 1  # Вставляем 1 в позициях ошибки

    # Генерация синдрома
    syndrome = generate_syndrome(error_vector, parity_check_matrix)
    syndrome_key = tuple(syndrome)

    # Добавляем в таблицу, если такого синдрома еще нет
    if syndrome_key not in syndrome_table:
        syndrome_table[syndrome_key] = error_vector

# Печать таблицы опознавателей
print("Таблица опознавателей для трёхкратных ошибок:")
for syndrome, error_vector in syndrome_table.items():
    print(f"Синдром: {syndrome}, Ошибочный вектор: {error_vector}")
