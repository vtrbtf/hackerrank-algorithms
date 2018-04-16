#!/usr/bin/env python3

x1, v1, x2, v2 = [int(n) for n in input('').split(' ')]
if (x1 > x2 and v1 >= v2) or (x2 > x1 and v2 >= v1):
    print("NO")
else:
    if (x1 - x2) % (v2 - v1) == 0:
        print ("YES")
    else:
        print ("NO")