import re

raw = open('./input.txt', 'r').read().split('\n')

def initialise(list, default):
    for i in range(1000):
        for j in range(1000):
            list[(i,j)] = default

def process(lights, input, turnon, toggle, turnoff):
    for inst in input:
        ints = re.findall(r'\d+', inst)

        for x in range(int(ints[0]), int(ints[2]) + 1):
                    for y in range(int(ints[1]), int(ints[3]) + 1):
                        if inst.startswith("turn on"):
                            lights[(x,y)] = turnon(lights[(x,y)])
                        elif inst.startswith("toggle"):
                            lights[(x,y)] = toggle(lights[(x,y)])
                        elif inst.startswith("turn off"):
                            lights[(x,y)] = turnoff(lights[(x,y)])

def part_one(raw):
    lights = {}
    on = 0

    initialise(lights, False)

    process(lights, raw, lambda x: True, lambda x: not x, lambda x: False)

    for light in lights:
        if lights[light]:
            on += 1
        
    return on

def part_two(raw):
    lights = {}
    lumens = 0

    initialise(lights, 0)

    process(lights, raw, lambda x: x + 1, lambda x: x + 2, lambda x: x - 1 if x > 0 else x)

    for light in lights:
        lumens += lights[light]

    return lumens

print(part_one(raw))
print(part_two(raw))