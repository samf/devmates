#! /usr/bin/env python3

from collections import defaultdict


def solver(occur):
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


def solve(s):
    d = defaultdict(int)
    for c in s:
        d[c] += 1
    occur = sorted(d.values())

    return solver(occur)


def demo():
    for i in ["eeeeffff", "aabbffddeaee", "llll", "", "aaabbbcccddddeeeee"]:
        print("input: ", i)
        print("output: ", solve(i))


if __name__ == "__main__":
    demo()
