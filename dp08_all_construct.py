def all_construct(target, word_bank):
    if target == "":
        return [[]]

    result = []

    for word in word_bank:
        if target.find(word) == 0:
            suffix = target[len(word):]
            suffixWays = all_construct(suffix, word_bank)
            targetWays = list(map(lambda w: [word,*w], suffixWays))
            result.extend(targetWays)

    return result


def main():
    print(all_construct('abcdef', ['ab', 'abc', 'cd', 'def', 'abcd', 'ef']))
    print(all_construct('', ['ab', 'abc', 'cd', 'def', 'abcd', 'ef']))
    print(all_construct('skateboard',  ['bo', 'rd', 'ate', 't', 'ska', 'sk', 'boar']))
    # print(all_construct('eeeeeeeeeeeeeeeeeeeeeeeeef',  ['e', 'ee', 'eee', 'eeee', 'eeeee', 'eeeeeee', 'eeeeeeeee']))

    # print(all_construct_dy('abcdef', ['ab', 'abc', 'cd', 'def', 'abcd'], {}))
    # print(all_construct_dy('skateboard',  ['bo', 'rd', 'ate', 't', 'ska', 'sk', 'boar'], {}))
    # print(all_construct_dy('eeeeeeeeeeeeeeeeeeeeeeeeeee',  ['e', 'ee', 'eee', 'eeee', 'eeeee', 'eeeeeee', 'eeeeeeee'], {}))

main()