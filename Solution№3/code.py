from collections import deque


def read_maze(filename):
    with open(filename, "r") as file:
        maze = [list(line.strip()) for line in file]
    return maze


def find_start_end(maze):
    start = end = None
    for i in range(len(maze)):
        for j in range(len(maze[i])):
            if maze[i][j] == 'S':
                start = (i, j)
            elif maze[i][j] == 'E':
                end = (i, j)
    return start, end


def bfs(maze, start, end):
    queue = deque()
    queue.append((start, [start]))
    visited = set()
    visited.add(start)

    while queue:
        (current, path) = queue.popleft()
        i, j = current

        if current == end:
            return path

        for di, dj in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
            ni, nj = i + di, j + dj
            if 0 <= ni < len(maze) and 0 <= nj < len(maze[0]) and maze[ni][nj] != '#' and (ni, nj) not in visited:
                visited.add((ni, nj))
                queue.append(((ni, nj), path + [(ni, nj)]))

    return None


def mark_path(maze, path):
    t = -1
    for (i, j) in path:
        t += 1
        if maze[i][j] not in ('S', 'E'):
            maze[i][j] = '*'
    return t

def print_maze(maze):
    return '\n'.join(''.join(row) for row in maze)


def write_output(filename, maze, steps):
    with open(filename, 'w') as file:
        file.write(f"Steps: {steps}\n")
        file.write(print_maze(maze) + '\n')

def main(input_filename, output_filename="Quest3-mazePython/output.txt"):
    maze = read_maze(input_filename)
    start, end = find_start_end(maze)
    if not start or not end:
        print("Не найдена начальная или конечная точка.")
        return

    path = bfs(maze, start, end)
    if path:
        steps = mark_path(maze, path)
        maze_str = print_maze(maze)
        print(f"Количество шагов: {steps}")
        print(maze_str)
        write_output(output_filename, maze, steps)
    else:
        print("Путь не найден.")
        with open(output_filename, 'w') as file:
            file.write("Путь не найден.\n")


if __name__ == "__main__":
    main("Quest3-mazePython/input.txt")
