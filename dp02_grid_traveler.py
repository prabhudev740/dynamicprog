def grid_traveler_dy(m, n, memo={}):
    if not (m and n):
        return 0

    if n == m == 1:
        return 1
    if m > n:
        key = f"{m} {n}"
    else:
        key = f"{n} {m}"

    if key in memo:
        return memo[key]
    
    memo[key] = grid_traveler_dy(m-1, n, memo) + grid_traveler_dy(m, n-1, memo)

    return  memo[key]
    
def main():
    # print(grid_traveler(1, 1))
    # print(grid_traveler(1, 0))
    # print(grid_traveler(0, 0))
    # print(grid_traveler(3, 2))
    # print(grid_traveler(13, 18))

    print(grid_traveler_dy(1, 1))
    print(grid_traveler_dy(4, 4))
    print(grid_traveler_dy(10, 10))
    print(grid_traveler_dy(15, 15))

main()