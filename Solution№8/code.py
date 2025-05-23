from string import ascii_uppercase

#Чтение файла и ввод числа
file = (open("input.txt").readline()).upper()

n = int(input())

#шифр Цезаря
def Cez(n, file):
    res = ''
    l =ascii_uppercase
    for i in file:
        res += l[(l.index(i) + n) % len(l)]
    return res
#шифр Атбаш
def Atbash(file):
    l = ascii_uppercase
    return file.translate(str.maketrans(l + l.upper(), l[::-1] + l.upper()[::-1]))

#Вывод результатов
print(Atbash(file))
print(Cez(n, file))
