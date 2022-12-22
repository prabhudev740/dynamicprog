def best_sum(target, nums):
    if target == 0:
        return []

    if target < 0:
        return

    shortest = None

    for val in nums:
        rem = target - val
        remRes = best_sum(rem, nums)
        if remRes is not None:
            remRes.append(val)
            if shortest is None or len(shortest) > len(remRes):
                shortest = remRes

    return shortest
        

def best_sum_dy(target, nums, memo={}):
    if target in memo:
        return memo[target]

    if target == 0:
        return []

    if target < 0:
        return

    shortest = None

    for val in nums:
        rem = target - val
        remRes = best_sum_dy(rem, nums, memo)
        if remRes is not None:
            combination = remRes + [val]

            if shortest is None or len(shortest) > len(combination):
                shortest = combination

    memo[target] = shortest
    return shortest

def main():
    # print(best_sum(7, [5, 3, 4, 7])) 
    # print(best_sum(8, [2, 3, 5])) 
    # print(best_sum(8, [1, 4, 5])) 
    # print(best_sum(100, [1, 2, 5, 25])) 

    print(best_sum_dy(7, [5, 3, 4, 7], {})) 
    print(best_sum_dy(8, [2, 3, 5], {})) 
    print(best_sum_dy(8, [1, 4, 5], {})) 
    print(best_sum_dy(9, [5], {})) 
    print(best_sum_dy(100, [1, 2, 5, 25], {})) 


main()