def load_matrix_data(file_path):
    with open(file_path, "r") as input_file:
        file_lines = input_file.readlines()
    matrix_data = []
    for line in file_lines:
        matrix_row = [float(num) for num in line.strip().split()]
        matrix_data.append(matrix_row)
    return matrix_data

def save_calculations(output_path, det_value, trace_value, transposed_matrix):
    with open(output_path, 'w') as output_file:
        output_file.write(f"Determinant: {det_value}\n")
        output_file.write(f"Trace: {trace_value}\n")
        output_file.write("Transposed Matrix:\n")
        for row in transposed_matrix:
            output_file.write(' '.join(map(str, row)) + '\n')

def calculate_matrix_trace(matrix):
    diagonal_sum = 0.0
    for idx in range(len(matrix)):
        diagonal_sum += matrix[idx][idx]
    return diagonal_sum

def create_transposed_matrix(original_matrix):
    transposed = []
    for col_idx in range(len(original_matrix[0])):
        transposed_row = []
        for row_idx in range(len(original_matrix)):
            transposed_row.append(original_matrix[row_idx][col_idx])
        transposed.append(transposed_row)
    return transposed

def create_submatrix(full_matrix, exclude_row, exclude_col):
    smaller_matrix = []
    for row_idx in range(len(full_matrix)):
        if row_idx == exclude_row:
            continue
        new_row = []
        for col_idx in range(len(full_matrix[row_idx])):
            if col_idx == exclude_col:
                continue
            new_row.append(full_matrix[row_idx][col_idx])
        smaller_matrix.append(new_row)
    return smaller_matrix

def calculate_matrix_determinant(matrix):
    matrix_size = len(matrix)
    if matrix_size == 1:
        return matrix[0][0]
    if matrix_size == 2:
        return matrix[0][0] * matrix[1][1] - matrix[0][1] * matrix[1][0]
    
    total_determinant = 0.0
    for column in range(matrix_size):
        submatrix = create_submatrix(matrix, 0, column)
        sign_multiplier = (-1) ** column
        submatrix_det = calculate_matrix_determinant(submatrix)
        total_determinant += sign_multiplier * matrix[0][column] * submatrix_det
    return total_determinant

def execute_matrix_operations():
    matrix = load_matrix_data("input.txt")
    
    row_count = len(matrix)
    for row in matrix:
        if len(row) != row_count:
            with open("output.txt", "w") as error_file:
                error_file.write("Error: Matrix must be square\n")
            return
    
    det = calculate_matrix_determinant(matrix)
    trace = calculate_matrix_trace(matrix)
    transposed = create_transposed_matrix(matrix)
    
    save_calculations("output.txt", det, trace, transposed)

if name == "main":
    execute_matrix_operations()
