#!/usr/bin/env python3

_n, _m = [int(n) for n in input('').split(' ')]
n = [int(j) for j in input('').split(' ')]
m = [int(j) for j in input('').split(' ')]

wholeset = n + m
counter = 0
for x in range(min(wholeset), max(wholeset) + 1):
    giveup = False

    for a in n:
        if x % a != 0:
            giveup = True
            break

    if giveup:
        continue

    for b in m:
        if b % x != 0:
            giveup = True
            break

    if not giveup:
        counter += 1


print(counter)
