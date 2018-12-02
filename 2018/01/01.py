#!/usr/bin/env python


def find_repeated_frequency(frequency_changes):

    frequency = 0
    visited = {}
    i = 0
    while(True):
        frequency += frequency_changes[i]
        if (visited.get(frequency) != None):
            return frequency

        visited[frequency] = True
        i = (i + 1) % len(frequency_changes)


def main():
    with open('input.txt', 'r') as file:
        frequency_changes = [int(line) for line in file]

    print('Part 1 answer = {}'.format(sum(frequency_changes)))
    print('Part 2 answer = {}'.format(
        find_repeated_frequency(frequency_changes)))


if __name__ == '__main__':
    main()
