#!/bin/python3

import math
import os
import random
import re
import sys

a = []

if __name__ == '__main__':
    N = int(input())

    for N_itr in range(N):
        firstNameEmailID = input().split()

        firstName = firstNameEmailID[0]

        emailID = firstNameEmailID[1]

        if re.search(".+@gmail\.com$", emailID):
            a.append(firstName)
    a.sort()

    for i in range(len(a)):
        print(a[i])
