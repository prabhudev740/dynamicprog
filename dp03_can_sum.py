def can_sum(target, nums):
    if not target:
        return True

    if target < 0:
        return False

    for val in nums:
        rem = target - val
        if can_sum(rem, nums):
            return True

    return False 

def can_sum_dy(target, nums, memo={}):
    if target in memo:
        return memo[target]

    if not target:
        return True

    if target < 0:
        return False

    for val in nums:
        rem = target - val
        if can_sum_dy(rem, nums, memo) == True:
            memo[target] = True
            return True

    memo[target] = False
    return False 

def main():
    # print(can_sum(7, [2, 3]))
    # print(can_sum(7, [5, 3, 4, 7]))
    # print(can_sum(7, [2, 4]))
    # print(can_sum(8, [2, 3, 5]))
    # print(can_sum(320, [7, 14]))

    print(can_sum_dy(7, [2, 3], {}))
    print(can_sum_dy(7, [5, 3, 4, 7], {}))
    print(can_sum_dy(7, [2, 4], {}))
    print(can_sum_dy(8, [2, 3, 5], {}))
    print(can_sum_dy(320, [7, 14], {}))
    
main()