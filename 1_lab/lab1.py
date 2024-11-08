import math
from heapq import heappop, heappush

# Заданные вероятности символов
probabilities = {
    'A': 0.204, 'B': 0.184, 'C': 0.176, 'D': 0.146, 
    'E': 0.134, 'F': 0.077, 'G': 0.071, 'H': 0.008
}

# Функция для построения кода Шеннона-Фано
def shannon_fano(symbol_probs):
    # Вспомогательная функция для рекурсивного построения кодов
    def recursive_split(symbols_probs, prefix=''):
        # Если остался один символ, присваиваем ему текущий префикс
        if len(symbols_probs) == 1:
            return {symbols_probs[0][0]: prefix}
        
        # Вычисляем суммарную вероятность символов
        total_prob = sum(prob for _, prob in symbols_probs)
        acc_prob = 0  # Накопленная вероятность для поиска места разделения
        split_idx = 0  # Индекс для разделения на две группы

        # Находим точку, где можно разделить символы на две группы с примерно равными вероятностями
        for i, (_, prob) in enumerate(symbols_probs):
            acc_prob += prob
            if acc_prob >= total_prob / 2:
                split_idx = i + 1
                break

        # Рекурсивно кодируем левую и правую части
        left_side = symbols_probs[:split_idx]
        right_side = symbols_probs[split_idx:]
        
        codes = {}
        # Добавляем '0' к префиксу для левой части и '1' для правой части
        codes.update(recursive_split(left_side, prefix + '0'))
        codes.update(recursive_split(right_side, prefix + '1'))
        return codes

    # Сортируем символы по вероятностям по убыванию для начала алгоритма
    sorted_symbols = sorted(symbol_probs.items(), key=lambda x: -x[1])
    return recursive_split(sorted_symbols)

# Функция для построения кода Хаффмана
def huffman(symbol_probs):
    # Инициализация кучи с парами (вероятность, символ)
    heap = [[prob, [sym, ""]] for sym, prob in symbol_probs.items()]
    # Строим дерево Хаффмана
    while len(heap) > 1:
        # Извлекаем два элемента с наименьшей вероятностью
        lo = heappop(heap)
        hi = heappop(heap)

        # Добавляем '0' к коду первого символа и '1' ко второму
        for pair in lo[1:]:
            pair[1] = '0' + pair[1]
        for pair in hi[1:]:
            pair[1] = '1' + pair[1]

        # Объединяем их в новый узел и добавляем обратно в кучу
        heappush(heap, [lo[0] + hi[0]] + lo[1:] + hi[1:])
    
    # Возвращаем словарь с кодами, отсортированный по длине кодов
    return dict(sorted(heappop(heap)[1:], key=lambda p: (len(p[-1]), p)))

# Функция для вычисления метрик кодирования
def calculate_metrics(symbol_probs, codes):
    # Средняя длина кода рассчитывается как сумма произведений вероятности символа и длины его кода
    avg_length = sum(symbol_probs[symbol] * len(code) for symbol, code in codes.items())

    # Энтропия источника: -∑(pi * log2(pi)), где pi — вероятность символа
    entropy = -sum(prob * math.log2(prob) for prob in symbol_probs.values())

    # Относительная эффективность: H / L, где H — энтропия, L — средняя длина кода
    efficiency = (entropy / avg_length) * 100  # В процентах

    # Коэффициент статистического сжатия: отношение фиксированной длины к средней длине
    fixed_length = math.ceil(math.log2(len(symbol_probs)))  # Фиксированная длина — минимальное целое от log2(N)
    compression_ratio = fixed_length / avg_length

    # Возвращаем все метрики
    return avg_length, entropy, efficiency, compression_ratio

# Шаги для построения кода и вычисления метрик для Шеннона-Фано
shannon_fano_codes = shannon_fano(probabilities)
shannon_fano_metrics = calculate_metrics(probabilities, shannon_fano_codes)

# Шаги для построения кода и вычисления метрик для Хаффмана
huffman_codes = huffman(probabilities)
huffman_metrics = calculate_metrics(probabilities, huffman_codes)

# Вывод результатов
print("Коды Шеннона-Фано:", shannon_fano_codes)
print("Метрики Шеннона-Фано (Средняя длина, Энтропия, Эффективность, Коэффициент сжатия):", shannon_fano_metrics)
print('')
print("Коды Хаффмана:", huffman_codes)
print("Метрики Хаффмана (Средняя длина, Энтропия, Эффективность, Коэффициент сжатия):", huffman_metrics)
