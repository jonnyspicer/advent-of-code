import re

input = open('./input.txt', 'r').read()

niceone, nicetwo = 0, 0
vowels = 'aeiou'

words = input.split('\n')

for word in words:
    if 'ab' in word:
        continue
    if 'cd' in word:
        continue
    if 'pq' in word:
        continue
    if 'xy' in word:
        continue

    v = 0
    for w in word:
        if w in vowels:
            v += 1  

    if v < 3:
        continue

    rx = re.compile(r'(.)\1{1,}')
    rxx = rx.search(word)
    if not rxx:
        continue

    niceone += 1

for word in words:
    arx = re.compile(r'(..).*?\1')
    arxx = arx.search(word)
    if not arxx:
        continue

    brx = re.compile(r'(.).\1')
    brxx = brx.search(word)
    if not brxx:
        continue

    nicetwo += 1

print(niceone)
print(nicetwo)