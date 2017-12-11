#! /usr/bin/env python3
import os
import sys

def main():
    if (len(sys.argv) > 1) :
        print("Part 1: " + str(solve_1(sys.argv[1])))
        print("Part 2: " + str(solve_2(sys.argv[1])))
    else:
        print("Part 1: " + str(solve_1(readInput())))
        print("Part 2: " + str(solve_2(readInput())))

def solve_1(str):
    sum = 0
    for i in range(len(str) - 1):
        if (str[i] == str[i + 1]):
            sum += int(str[i])

    if (str[len(str) -1] == str[0]):
        sum += int(str[0])

    return sum

def solve_2(str):
    str_len = len(str)
    sum = 0
    for i in range(str_len):
        if (str[i] == str[(i + str_len // 2) % str_len]):
            sum += int(str[i])

    return sum


def readInput():
    return open(os.path.join(sys.path[0], '01.txt'), 'r').read()

if __name__ == '__main__':
    main()