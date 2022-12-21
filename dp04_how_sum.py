def how_sum(target, nums):
    if target == 0:
        return []
    
    if target < 0:
        return None

    for val in nums:
        rem = target - val

        remRes = how_sum(rem, nums)
        if remRes is not None:
            return remRes + [val]

    return None

def how_sum_dy(target, nums, memo={}):
    if target in memo:
        return memo[target]


    if target == 0:
        return []

    if target < 0:
        return None

    for val in nums:
        rem = target - val
        remRes = how_sum_dy(rem, nums, memo)
        
        if remRes is not None:
            remRes.append(val)
            memo[target] = remRes
            return memo[target]

    memo[target] = None
    return None


def main():
    # print(how_sum(7, [2, 3]))
    # print(how_sum(7, [5, 3, 4, 7]))
    # print(how_sum(7, [2, 4]))
    # print(how_sum(8, [2, 3, 5]))
    # print(how_sum(0, [2, 3, 5]))
    # print(how_sum(320, [7, 14]))

    print(how_sum_dy(7, [2, 3]), {})
    print(how_sum_dy(7, [5, 3, 4, 7], {}))
    print(how_sum_dy(7, [2, 4], {}))
    print(how_sum_dy(8, [2, 3, 5], {}))
    print(how_sum_dy(0, [2, 3, 5], {}))
    print(how_sum_dy(320, [7, 14], {}))

main()