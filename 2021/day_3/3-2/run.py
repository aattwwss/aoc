from copy import copy
def loadInput(filename):
    file = open(filename, 'r')
    lines = file.readlines()
    return [line.strip() for line in lines]

    

lines = loadInput("./input.txt")

oxygen = copy(lines)
for i in range(12):
    if len(oxygen) == 1:
        break
    currentBits = [x[i] for x in oxygen]
    count1 = currentBits.count("1")
    count0 = currentBits.count("0")
    winningBit = "1"
    if count1 < count0:
        winningBit = "0"
    oxygen = [x for x in oxygen if x[i] == winningBit]
co2 = copy(lines)
for i in range(12):
    if len(co2) == 1:
        break
    currentBits = [x[i] for x in co2]
    count1 = currentBits.count("1")
    count0 = currentBits.count("0")
    winningBit = "0"
    if count1 < count0:
        winningBit = "1"
    co2 = [x for x in co2 if x[i] == winningBit]

print(oxygen, co2)
print(int(oxygen[0],2)*int(co2[0],2))
