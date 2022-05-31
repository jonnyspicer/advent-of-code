input = open('./input.txt', 'r').read()
total = 0
ribbon = 0

prisms = input.split('\n')

# for each prism
# calculate 2*l*w, 2*w*h, 2*h*l
# calculate smallest side
# add them all up

for prism in prisms:
    feet = 0
    sides = prism.split('x')
    for i in range(0, len(sides)):
        sides[i] = int(sides[i])

    lw = 2 * sides[0] * sides[1]
    hw = 2 * sides[1] * sides[2]
    hl = 2 * sides[0] * sides[2]
    smallest = min([lw, hw, hl])
    ribbon += (sides[0] * sides[1] * sides[2])
    ribbon += (2 * sides[0])
    ribbon += (2 * sides[1])
    ribbon += (2 * sides[2])
    ribbon -= (2 * max([sides[0], sides[1], sides[2]]))
    total += (lw + hw + hl + (smallest / 2))

print(ribbon)
print(total)
