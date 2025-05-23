#include <iostream>
#include <fstream>
using namespace std;

//Чтение матрицы
void readMatrix(ifstream& inFile, int**& matrix, int& rows, int& cols) {
    inFile >> rows >> cols;
    matrix = new int* [rows];
    for (int i = 0; i < rows; i++) {
        matrix[i] = new int[cols];
        for (int j = 0; j < cols; j++) {
            inFile >> matrix[i][j];
        }
    }
}

//Умножение матриц
void multiplyMatrices(int** matrix1, int** matrix2, int**& result, int rows1, int cols1, int rows2, int cols2) {
    if (cols1 != rows2) {
        cerr << "n/a" << endl;
        return;
    }

    result = new int* [rows1];
    for (int i = 0; i < rows1; i++) {
        result[i] = new int[cols2];
        for (int j = 0; j < cols2; j++) {
            result[i][j] = 0;
            for (int k = 0; k < cols1; k++) {
                result[i][j] += matrix1[i][k] * matrix2[k][j];
            }
        }
    }
}

//Запись матрицы в файл
void writeMatrix(ofstream& outFile, int** matrix, int rows, int cols) {
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            outFile << matrix[i][j] << " ";
        }
        outFile << endl;
    }
}

//Освобождение памяти
void freeMatrix(int**& matrix, int rows) {
    for (int i = 0; i < rows; i++) {
        delete[] matrix[i];
    }
    delete[] matrix;
    matrix = nullptr;
}

//Главная функция
int main() {
    ifstream inFile("input.txt");
    ofstream outFile("output.txt");

    if (!inFile.is_open() || !outFile.is_open()) {
        cerr << "n/a" << endl;
        return 1;
    }

    int rows1, cols1, rows2, cols2;
    int** matrix1 = nullptr, ** matrix2 = nullptr, ** result = nullptr;

    readMatrix(inFile, matrix1, rows1, cols1);

    readMatrix(inFile, matrix2, rows2, cols2);

    if (cols1 != rows2) {
        cerr << "n/a" << endl;
        freeMatrix(matrix1, rows1);
        freeMatrix(matrix2, rows2);
        return 1;
    }

    multiplyMatrices(matrix1, matrix2, result, rows1, cols1, rows2, cols2);

    writeMatrix(outFile, result, rows1, cols2);

    freeMatrix(matrix1, rows1);
    freeMatrix(matrix2, rows2);
    freeMatrix(result, rows1);

    inFile.close();
    outFile.close();

    cout << "Done" << endl;
    return 0;
}
