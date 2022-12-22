def can_construct(target, word_bank):
    if target == "":
        return True

    for word in word_bank:
        if target.find(word) == 0:
            suffix = target[len(word):]
            if can_construct(suffix, word_bank):
                return True

    return False

def can_construct_dy(target, word_bank, memo):
    if target in memo:
        return memo[target]

    if target == "":
        return True

    for word in word_bank:
        if target.find(word) == 0:
            suffix = target[len(word):]
            if can_construct_dy(suffix, word_bank, memo):
                memo[target] = True
                return True
    memo[target] = False
    return False


def main():
    print(can_construct('abcdef', ['ab', 'abc', 'cd', 'def', 'abcd']))
    print(can_construct('skateboard',  ['bo', 'rd', 'ate', 't', 'ska', 'sk', 'boar']))
    print(can_construct('eeeeeeeeeeeeeeeeeeeeeeeeef',  ['e', 'ee', 'eee', 'eeee', 'eeeee', 'eeeeeee', 'eeeeeeee']))
    
    print(can_construct_dy('abcdef', ['ab', 'abc', 'cd', 'def', 'abcd'], {}))
    print(can_construct_dy('skateboard',  ['bo', 'rd', 'ate', 't', 'ska', 'sk', 'boar'], {}))
    print(can_construct_dy('eeeeeeeeeeeeeeeeeeeeeeeeeeeeeef',  ['e', 'ee', 'eee', 'eeee', 'eeeee', 'eeeeeee', 'eeeeeeee'], {}))

main()