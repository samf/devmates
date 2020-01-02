#! /usr/bin/env python

from collections import defaultdict

"""
Given a string s consisting of n lowercase letters, you have to
delete the minimum number of characters from s so that every letter
in s appears a unique number of times. We only care about the
occurrences of letters that appear at least once in result.

 Input: "eeeeffff"
 Output: 1
 Why? We can delete one occurence of e or one occurence of 'f'.
       Then one letter will occur four times and the other three times.

 Input: "aabbffddeaee"
 Output:  6
 Why:
   For example, we can delete all occurences of 'e' and 'f'
   and one occurence of 'd' to obtain the word "aabbda".
   Note that both 'e' and 'f' will occur zero times in the new word,
   but that's fine, since we only care about the letter
   that appear at least once.

 Input: "llll"
 Output: 0
 Why? There is no need to delete any character.
"""


def solve(occur):
    rc = 0
    done = False

    while not done:
        for i, left, right in zip(range(len(occur)), occur, occur[1:]):
            if left == 0:
                continue
            if left == right:
                occur[i] -= 1
                rc += 1
                break
        else:
            done = True

    return rc


def collate(s):
    d = defaultdict(int)
    for c in s:
        d[c] += 1
    return sorted(d.values())


def demo():
    for input in ["eeeeffff", "aabbffddeaee", "llll", "", "aaabbbcccddddeeeee"]:
        print("input: ", input)
        print("output: ", solve(collate(input)))


if __name__ == "__main__":
    demo()
