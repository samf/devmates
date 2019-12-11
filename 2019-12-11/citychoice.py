#! /usr/bin/env python3

"""
There are 2N people a company is planning to interview. The cost
of flying the i-th person to city A is costs[i][0], and the cost
of flying the i-th person to city B is costs[i][1].

Return the minimum cost to fly every person to a city such that
exactly N people arrive in each city.

  Input: [[10,20],[30,200],[400,50],[30,20]]
  Output: 110
  Why?
  The first person goes to city A for a cost of 10.
  The second person goes to city A for a cost of 30.
  The third person goes to city B for a cost of 50.
  The fourth person goes to city B for a cost of 20.

  The total minimum cost is 10 + 30 + 50 + 20 = 110
  to have half the people interviewing in each city.

  Input: [[259,770],[448,54],[926,667],[184,139],[840,118],[577,469]]
  Output: 1859
"""


class Plan:
    def __init__(self, cost_array, expected=None):
        self.people = [Person(*i) for i in cost_array]
        self.expected = expected
        self.city0 = list()
        self.city1 = list()
        self.total = 0

    def solve(self):
        for city, base, upgrade in [p.naive() for p in self.people]:
            self.total += base
            if city == 0:
                self.city0.append(upgrade)
            else:
                self.city1.append(upgrade)

        self.city0.sort()
        self.city1.sort()

        city = self.city0
        imbalance = len(self.city0) - len(self.city1)
        if imbalance < 0:
            imbalance *= -1
            city = self.city1

        while imbalance > 0:
            self.total += city.pop(0)
            imbalance -= 2

class Person:
    def __init__(self, city0, city1):
        self.city0 = city0
        self.city1 = city1
        self.upgrade = 0

    def naive(self):
        "return city choice, cost, and upgrade cost"
        if self.city0 > self.city1:
            return 1, self.city1, self.city0 - self.city1
        return 0, self.city0, self.city1 - self.city0


def main():
    for case in [
        Plan([[10, 20], [30, 200], [400, 50], [30, 20]], 110),
        Plan(
            [
                [259, 770],
                [448, 54],
                [926, 667],
                [184, 139],
                [840, 118],
                [577, 469],
            ],
            1859,
        ),
    ]:
        case.solve()
        if case.expected and case.total != case.expected:
            print(f"expected {case.expected} but got {case.total}")
        else:
            print(case.total)


if __name__ == "__main__":
    main()
