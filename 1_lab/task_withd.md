
### Пояснения по коду:
- shannon_fano:
Сначала сортирует символы по убыванию вероятности.
Разделяет упорядоченный список на две группы, чтобы суммы вероятностей в каждой были максимально близки.
Каждой группе присваивается бит ('0' или '1'), и для каждой группы процесс повторяется рекурсивно.

- huffman:
Использует кучу (min-heap) для объединения двух наименее вероятных символов в составной символ.
Присваивает '0' и '1' в процессе объединения, пока не останется один составной символ.
Код символов строится по дереву, начиная с листьев к корню.

- calculate_metrics:
Средняя длина кода рассчитывается как ожидаемое значение длины кодового слова.
Энтропия измеряет теоретическую минимальную длину для кодирования символов на основе их вероятностей.
Эффективность указывает, насколько эффективен построенный код относительно энтропии источника.
Коэффициент сжатия сравнивает кодовую длину с фиксированной длиной, чтобы понять, насколько код более эффективен по сравнению с равномерным кодированием.

#### Числа для решения 

0.204 
0.184
0.176
0.146
0.134
0.077
0.071
0.008

0.283
0.275
0.215
0.072
0.057
0.049
0.033
0.016

#### Обьяснение 
Этот код создаёт оптимальные неравномерные коды для символов с заданными вероятностями появления. Коды строятся с использованием двух методов сжатия информации: метода Шеннона-Фано и метода Хаффмана. После построения кодов код рассчитывает метрики, которые помогают оценить качество этих кодов.

**Входные данные**

_На вход поступает список символов с их вероятностями. В нашем случае это следующие вероятности:_


Каждое значение SymbolProb содержит:
Symbol — имя символа (например, "A").
Probability — вероятность его появления в сообщении (например, 0.204).

*Логика кода*

Метод Шеннона-Фано (функция ShannonFano)
Эта функция создает коды для символов с помощью метода Шеннона-Фано.

Она сортирует символы по убыванию вероятности и делит их на две группы с примерно одинаковыми суммами вероятностей.
Каждой группе присваивается бит "0" или "1", и затем происходит рекурсивное деление на подгруппы.

В результате каждому символу назначается уникальный код, построенный из нулей и единиц. Эти коды выводятся как словарь вида {Symbol: Code}, например: {"A": "0", "B": "10", ...}.

**Метод Хаффмана (функция Huffman)**
Метод Хаффмана также строит коды для символов, но с использованием "кучи" (структуры данных приоритетной очереди).
Сначала символы с меньшими вероятностями объединяются, образуя составные узлы дерева, каждый из которых обозначается битом "0" или "1".
Когда дерево завершено, каждому символу присваивается уникальный путь от корня до листа, который и является его кодом.

Эти коды также выводятся как словарь {Symbol: Code}, аналогично коду Шеннона-Фано.
Расчёт метрик кодирования (функция CalculateMetrics)

Функция вычисляет четыре ключевые метрики, чтобы оценить качество кодов:
- Средняя длина кода — усреднённое количество битов на символ в коде. Это значение зависит от вероятностей символов и их длин в закодированном сообщении.
- Энтропия источника — теоретическая минимальная длина кода (в битах на символ), вычисленная на основе вероятностей символов. Она указывает, насколько сжатым может быть сообщение.
- Эффективность — показывает, насколько эффективно кодирование по сравнению с энтропией источника. Если эффективность равна 100%, значит, код близок к теоретически минимально возможной длине.
- Коэффициент сжатия — отношение фиксированной длины (код с одинаковой длиной символов) к средней длине нашего кода. Чем больше этот коэффициент, тем лучше сжатие по сравнению с кодом фиксированной длины.


#### Вывод
Средняя длина для метода Хаффмана, как правило, меньше, чем для Шеннона-Фано, так как код Хаффмана обеспечивает оптимальную длину для данного набора вероятностей.

Энтропия показывает теоретическую нижнюю границу длины. Обе методики обычно дают среднюю длину кода, близкую к энтропии.

Эффективность для Хаффмана часто близка к 100%, что делает его оптимальным. Эффективность для Шеннона-Фано также может быть высокой, но чуть ниже.

Коэффициент сжатия: у Хаффмана он часто выше, так как он использует меньшее количество битов.

Уменьшение алфавита ведёт к более коротким кодам, меньшей энтропии и высокой эффективности.
Увеличение алфавита приводит к увеличению средней длины кода и энтропии, а эффективность может немного снизиться.