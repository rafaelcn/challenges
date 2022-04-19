x = int(input())

left = 1
right = 10**18

ans = 0

while left <= right:
    m = (left + right) // 2

    sum = m * (m + 1) // 2

    if sum > x:
        right = m - 1
        ans = m
    else:
        left = m + 1

print(ans)
