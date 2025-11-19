count = int(input().strip())
input1 = input().strip()
input2 = input().strip()

errorCount = 0

for i in range(count):
    if input1[i] != input2[i]:
        errorCount += 1

print(errorCount)
