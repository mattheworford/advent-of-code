def score(moves):
    if moves[0] == 'A':
        if moves[1] == 'X':
            return 3
        elif moves[1] == 'Y':
            return 4
        elif moves[1] == 'Z':
            return 8
    elif moves[0] == 'B':
        if moves[1] == 'X':
            return 1
        elif moves[1] == 'Y':
            return 5
        elif moves[1] == 'Z':
            return 9
    elif moves[0] == 'C':
        if moves[1] == 'X':
            return 2
        elif moves[1] == 'Y':
            return 6
        elif moves[1] == 'Z':
            return 7


with open('input.txt', 'r') as input_file:
    total = 0
    for line in input_file:
        moves = line.strip().split(' ')
        total += score(moves)
    print(total)
