from string import ascii_uppercase

# Чтение входных данных
input_text = (open("input.txt").readline()).upper()
shift_value = int(input())

# Шифр Цезаря
def apply_caesar_cipher(shift, text):
    encrypted_text = ''
    alphabet = ascii_uppercase
    for char in text:
        if char in alphabet:
            new_position = (alphabet.index(char) + shift) % len(alphabet)
            encrypted_text += alphabet[new_position]
        else:
            encrypted_text += char
    return encrypted_text

# Шифр Атбаш
def apply_atbash_cipher(text):
    alphabet = ascii_uppercase
    translation_table = str.maketrans(
        alphabet + alphabet.lower(),
        alphabet[::-1] + alphabet.lower()[::-1]
    )
    return text.translate(translation_table)

# Обработка и вывод результатов
encrypted_atbash = apply_atbash_cipher(input_text)
encrypted_caesar = apply_caesar_cipher(shift_value, input_text)

print(encrypted_atbash)
print(encrypted_caesar)
