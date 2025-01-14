input = open("input.txt")

res = 0

for line in input:

    temp = ""
    
    start, end = 0, len(line) - 1

    while start < end:
        if line[start].isnumeric():
            
    