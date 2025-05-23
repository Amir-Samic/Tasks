#include <fstream>
#include <vector>
#include <iostream>

using namespace std;

// Функция загрузки матрицы из файла
vector<vector<double>> loadMatrixFromFile() {
    ifstream inputFile("input.txt");
    vector<vector<double>> matrixData;

    if (!inputFile.is_open()) {
        cout << "Ошибка открытия файла!" << endl;
        return matrixData;
    }

    double value;
    while (inputFile >> value) {
        vector<double> currentRow;
        currentRow.push_back(value);

        while (inputFile.peek() != '\n' && inputFile.peek() != EOF) {
            inputFile >> value;
            currentRow.push_back(value);
        }

        matrixData.push_back(currentRow);
    }

    inputFile.close();
    return matrixData;
}

// Вычисление следа матрицы
double computeMatrixTrace(const vector<vector<double>>& matrix) {
    double traceValue = 0;
    for (int i = 0; i < matrix.size(); i++) {
        traceValue += matrix[i][i];
    }
    return traceValue;
}

// Транспонирование матрицы
vector<vector<double>> computeMatrixTranspose(const vector<vector<double>>& matrix) {
    vector<vector<double>> transposedMatrix(matrix[0].size(), vector<double>(matrix.size()));

    for (int row = 0; row < matrix.size(); row++) {
        for (int col = 0; col < matrix[0].size(); col++) {
            transposedMatrix[col][row] = matrix[row][col];
        }
    }

    return transposedMatrix;
}

// Рекурсивное вычисление определителя
double computeMatrixDeterminant(vector<vector<double>> matrix) {
    int matrixSize = matrix.size();

    if (matrixSize == 1) return matrix[0][0];
    if (matrixSize == 2) {
        return matrix[0][0] * matrix[1][1] - matrix[0][1] * matrix[1][0];
    }

    double determinantValue = 0;

    for (int col = 0; col < matrixSize; col++) {
        vector<vector<double>> minorMatrix(matrixSize - 1, vector<double>(matrixSize - 1));

        for (int row = 1; row < matrixSize; row++) {
            int minorCol = 0;
            for (int matrixCol = 0; matrixCol < matrixSize; matrixCol++) {
                if (matrixCol == col) continue;
                minorMatrix[row - 1][minorCol] = matrix[row][matrixCol];
                minorCol++;
            }
        }

        double signCoefficient = (col % 2 == 0) ? 1 : -1;
        determinantValue += signCoefficient * matrix[0][col] * computeMatrixDeterminant(minorMatrix);
    }

    return determinantValue;
}

// Сохранение результатов в файл
void saveResultsToFile(double traceValue, double determinantValue, 
                      const vector<vector<double>>& transposedMatrix) {
    ofstream outputFile("output.txt");

    outputFile << "След матрицы: " << traceValue << endl;
    outputFile << "Определитель матрицы: " << determinantValue << endl;
    outputFile << "Транспонированная матрица:" << endl;

    for (const auto& row : transposedMatrix) {
        for (double value : row) {
            outputFile << value << " ";
        }
        outputFile << endl;
    }

    outputFile.close();
}

int main() {
    vector<vector<double>> matrix = loadMatrixFromFile();

    if (matrix.empty() || matrix.size() != matrix[0].size()) {
        cout << "Ошибка: матрица должна быть квадратной" << endl;
        return 1;
    }

    double trace = computeMatrixTrace(matrix);
    double determinant = computeMatrixDeterminant(matrix);
    vector<vector<double>> transposed = computeMatrixTranspose(matrix);

    saveResultsToFile(trace, determinant, transposed);

    cout << "Результаты сохранены в output.txt" << endl;

    return 0;
}
