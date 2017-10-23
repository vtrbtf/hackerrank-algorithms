#!/usr/bin/env python3

house = [int(n) for n in input('input house coords \n').split(' ')]
trees = [int(n) for n in input('Input trees coords \n').split(' ')]
m_and_n = [int(n) for n in input('input m and n coords \n').split(' ')]
m_distances = [int(n) for n in input('input m dist \n').split(' ')]
n_distances = [int(n) for n in input('input n dist \n').split(' ')]

m_counter = 0
for md in m_distances:
    if md > 0:
        ground_dist = trees[0] + md
        if ground_dist >= house[0] and ground_dist <= house[1]:
            m_counter += 1


n_counter = 0
for nd in n_distances:
    if nd < 0:
        ground_dist = trees[1] + nd
        if ground_dist >= house[0] and ground_dist <= house[1]:
            n_counter += 1

print(m_counter)
print(n_counter)
