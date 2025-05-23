from collections import Counter

# Чтение чисел из файлов
with open('A.txt', 'r') as f:
    numbers_a = list(map(int, f.read().split()))

with open('B.txt', 'r') as f:
    numbers_b = list(map(int, f.read().split()))

#Находим пересечение множеств
set_a = set(numbers_a)
set_b = set(numbers_b)
intersection = set_a & set_b

#Подсчитываем вхождения в исходных файлах
count_a = Counter(numbers_a)
count_b = Counter(numbers_b)

#Для каждого числа в пересечении находим максимальное количество вхождений
result = []
for num in sorted(intersection):
    max_count = max(count_a[num], count_b[num])
    result.extend([num] * max_count)

#Записываем результат в файл
with open('C.txt', 'w') as f:
    f.write(' '.join(map(str, result)))
