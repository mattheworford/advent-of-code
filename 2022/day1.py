import heapq

curr_calories = 0
heap = []

with open('input.txt', 'r') as input_file:
    for line in input_file:
        if line == '\n':
            if len(heap) < 3:
                heapq.heappush(heap, curr_calories)
            else:
                min_max = heapq.heappop(heap)
                heapq.heappush(heap, max(min_max, curr_calories))
            curr_calories = 0
        else:
            curr_calories += int(line)

print(heap)
print(sum(heap))
