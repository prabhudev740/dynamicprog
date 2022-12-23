def count_construct(target, word_banks):
    if target == "":
        return 1

    count = 0

    for word in word_banks:
        if target.find(word) == 0:
            suffix = target[len(word):]
            no_of_ways = count_construct(suffix, word_banks)
            count += no_of_ways

    return count


def count_construct_dy(target, word_banks, memo):
    if target in memo:
        return memo[target]

    if target == "":
        return 1

    count = 0

    for word in word_banks:
        if target.find(word) == 0:
            suffix = target[len(word):]
            no_of_ways = count_construct_dy(suffix, word_banks, memo)
            count += no_of_ways

    memo[target] = count
    return count

def main():
    # print(count_construct('abcdef', ['ab', 'abc', 'cd', 'def', 'abcd']))
    # print(count_construct('skateboard',  ['bo', 'rd', 'ate', 't', 'ska', 'sk', 'boar']))
    # print(count_construct('eeeeeeeeeeeeeeeeeeeeeeeeef',  ['e', 'ee', 'eee', 'eeee', 'eeeee', 'eeeeeee', 'eeeeeeee']))
    
    print(count_construct_dy('abcdef', ['ab', 'abc', 'cd', 'def', 'abcd'], {}))
    print(count_construct_dy('skateboard',  ['bo', 'rd', 'ate', 't', 'ska', 'sk', 'boar'], {}))
    print(count_construct_dy('eeeeeeeeeeeeeeeeeeeeeeeeeee',  ['e', 'ee', 'eee', 'eeee', 'eeeee', 'eeeeeee', 'eeeeeeee'], {}))

main()