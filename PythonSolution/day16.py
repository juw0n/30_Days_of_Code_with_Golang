#!/bin/python3

import sys

S = input().strip()
try:
    newEntry = int(S)
    print(newEntry)
except:
    print('Bad String')