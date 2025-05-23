import copy

def read_matrix(filename):
    with open(filename, "r") as file:
        lines = file.readlines()
    matrix = []
    for line in lines:
        row = [float(x) for x in line.strip().split()]
        matrix.append(row)
    return matrix

def write_results(filename, determinant, trace, transposed):
    with open(filename, 'w') as file:
        file.write(f"Determiant: {determinant}\n")
        file.write(f"Trace: {trace}\n")
        file.write("Transposed:\n")
        for row in transposed:
            file.write(' '.join(map(str, row)) + '\n')

def get_trace(matrix):
    trace = 0.0
    for i in range(len(matrix)):
        trace += matrix[i][i]
    return trace

def transpose_matrix(matrix):
    transposed = []
    for j in range(len(matrix[0])):
        new_row = []
        for i in range(len(matrix)):
            new_row.append(matrix[i][j])
        transposed.append(new_row)
    return transposed

def get_submatrix(matrix, row_to_remove, col_to_remove):
    submatrix = []
    for i in range(len(matrix)):
        if i == row_to_remove:
            continue
        new_row = []
        for j in range(len(matrix[i])):
            if j == col_to_remove:
                continue
            new_row.append(matrix[i][j])
        submatrix.append(new_row)
    return submatrix

def compute_determinant(matrix):
    n = len(matrix)
    if n == 1:
        return matrix[0][0]
    if n == 2:
        return matrix[0][0] * matrix[1][1] - matrix[0][1] * matrix[1][0]
    
    determinant = 0.0
    for j in range(n):
        submatrix = get_submatrix(matrix, 0, j)
        sign = (-1) ** j
        sub_det = compute_determinant(submatrix)
        determinant += sign * matrix[0][j] * sub_det
    return determinant

def main():
    matrix = read_matrix("input.txt")
    
    rows = len(matrix)
    for row in matrix:
        if len(row) != rows:
            with open("output.txt", "w") as file:
                file.write("Ошибка: матрица не квадратная\n")
            return
    
    determinant = compute_determinant(matrix)
    trace = get_trace(matrix)
    transposed = transpose_matrix(matrix)
    
    write_results("output.txt", determinant, trace, transposed)

if __name__ == "__main__":
    main()
