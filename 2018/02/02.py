def part_1(box_ids):
    two_count = 0
    three_count = 0

    for box_id in box_ids:
        if any([box_id.count(char) == 2 for char in set(box_id)]):
            two_count += 1
        if any([box_id.count(char) == 3 for char in set(box_id)]):
            three_count += 1
        
    return two_count * three_count


def part_2(box_ids):

    for box_id_1 in box_ids:
        for box_id_2 in box_ids:
            matching_chars = ''
            for char_1, char_2 in zip(box_id_1, box_id_2):
                if char_1 is char_2:
                    matching_chars += char_1
    
            if (len(box_id_1) - 1) == len(matching_chars):
                return matching_chars
def main():
    with open('input.txt', 'r') as file:
        box_ids = [line.strip() for line in file]

    print('Part 1 answer = {}'.format(part_1(box_ids)))
    print('Part 2 answer = {}'.format(part_2(box_ids)))


if __name__ == '__main__':
    main()
