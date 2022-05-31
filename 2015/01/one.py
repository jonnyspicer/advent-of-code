floor = 0
count = 0

input = open('./input.txt', 'r').read().strip()

for direction in input:
    count += 1
    if direction == ')':
        floor -= 1
    if direction == '(':
        floor += 1
    if floor < 0:
        print(count)

print(floor)
