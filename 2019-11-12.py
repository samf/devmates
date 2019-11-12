#! /usr/bin/env python3

from collections import defaultdict


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
