import random
from collections import Counter

#Генерация случайных чисел и запись в файлы
random.seed(42)

numbers_a = [random.randint(0, 500) for _ in range(1000)]
with open('A.txt', 'w') as f:
    f.write(' '.join(map(str, numbers_a)))

numbers_b = [random.randint(200, 700) for _ in range(1000)]
with open('B.txt', 'w') as f:
    f.write(' '.join(map(str, numbers_b)))

#Чтение чисел из файлов
with open('A.txt', 'r') as f:
    numbers_a = list(map(int, f.read().split()))

with open('B.txt', 'r') as f:
    numbers_b = list(map(int, f.read().split()))

#Находим пересечение множеств
set_a = set(numbers_a)
set_b = set(numbers_b)
intersection = sorted(set_a & set_b)  # Сразу сортируем

#Подсчитываем вхождения в исходных файлах
count_a = Counter(numbers_a)
count_b = Counter(numbers_b)

#Формируем результат с повторениями
result = []
for num in intersection:
    max_count = max(count_a[num], count_b[num])
    result.extend([num] * max_count)

#Записываем результат в файл
with open('C.txt', 'w') as f:
    for i in range(0, len(result), 10):
        f.write(' '.join(map(str, result[i:i+10])) + '\n')
