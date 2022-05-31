seen = []

input = open('./input.txt', 'r').read().strip()

santa = (0, 0)
robosanta = (0,0)
seen.append(santa)

for i in range(0, len(input)):
    if i % 2 == 0:
        if input[i] == '>':
            santa = (santa[0] + 1, santa[1])
        elif input[i] == '^':
            santa = (santa[0], santa[1] + 1)
        elif input[i] == '<':
            santa = (santa[0] - 1, santa[1])
        elif input[i] == 'v':
            santa = (santa[0], santa[1] - 1)
    else:
        if input[i] == '>':
            robosanta = (robosanta[0] + 1, robosanta[1])
        elif input[i] == '^':
            robosanta = (robosanta[0], robosanta[1] + 1)
        elif input[i] == '<':
            robosanta = (robosanta[0] - 1, robosanta[1])
        elif input[i] == 'v':
            robosanta = (robosanta[0], robosanta[1] - 1)

    if santa not in seen:
        seen.append(santa)
    if robosanta not in seen:
        seen.append(robosanta)

print(len(seen))