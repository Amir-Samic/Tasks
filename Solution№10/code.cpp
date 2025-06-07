#include <iostream>
#include <fstream>
#include <string>
#include <cctype>

using namespace std;

// Функция шифра Атбаш
string atbash(const string& text) {
    string result;
    for (char c : text) {
        if (isupper(c)) {
            result += 'Z' - (c - 'A');
        } else if (islower(c)) {
            result += 'z' - (c - 'a');
        } else {
            result += c;
        }
    }
    return result;
}

// Функция шифра Цезаря
string caesar(const string& text, int shift) {
    string result;
    for (char c : text) {
        if (isupper(c)) {
            result += ((c - 'A' + shift) % 26) + 'A';
        } else if (islower(c)) {
            result += ((c - 'a' + shift) % 26) + 'a';
        } else {
            result += c;
        }
    }
    return result;
}

int main() {
    // Чтение из файла
    ifstream inFile("input.txt");
    if (!inFile) {
        cerr << "Ошибка открытия файла input.txt" << endl;
        return 1;
    }

    string inputText((istreambuf_iterator<char>(inFile)), istreambuf_iterator<char>());

    // Ввод сдвига
    int shift;
    cout << "Введите сдвиг для шифра Цезаря: ";
    cin >> shift;

    // Шифрование
    string atbashText = atbash(inputText);
    string caesarText = caesar(inputText, shift);

    // Запись в файл
    ofstream outFile("output.txt");
    if (!outFile) {
        cerr << "Ошибка создания файла output.txt" << endl;
        return 1;
    }

    outFile << "Атбаш: " << atbashText << endl;
    outFile << "Цезарь (сдвиг " << shift << "): " << caesarText << endl;
    outFile.close();

    cout << "Готово! Результат записан в output.txt" << endl;

    return 0;
}
