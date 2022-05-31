import hashlib

input = open('./input.txt', 'r').read().strip()
i = 0

while i >= 0:
    inp = input + str(i)
    if hashlib.md5(inp.encode()).hexdigest().startswith('000000'):
        print(i)
        break
    i += 1
