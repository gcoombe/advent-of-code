def part_1(box_ids):
    two_count = 0
    three_count = 0

    for box_id in box_ids:
        sorted_str = ''.join(sorted(box_id))

        two_found = False
        three_found = False
        match_count = 1
        for i in range(1, len(sorted_str)):
        
            if (sorted_str[i] is sorted_str[i - 1]):
                match_count += 1
            else:
                if match_count == 2:
                    two_found = True
                elif match_count == 3:
                    three_found = True
                
                match_count = 1
        if match_count == 2:
            two_found = True
        elif match_count == 3:
            three_found = True

        
        if (two_found):
            two_count += 1
        if (three_found):
            three_count += 1
    
    return two_count * three_count

def check_strings(str_1, str_2):
    varying_char = None
    for j in range(0, len(str_1)):
        if str_1[j] is not str_2[j]:
            if varying_char != None:
                return None
            varying_char = str_1[j]
    if (varying_char != None):
        return str_1.translate(None, varying_char)
    else:
        return None

def part_2(box_ids):

    for i in range(0, len(box_ids)):
        for k in range(i, len(box_ids)):
            found_str = check_strings(box_ids[i], box_ids[k])

            if (found_str != None):
                return found_str
def main():
    with open('input.txt', 'r') as file:
        box_ids = [line.strip() for line in file]

    print('Part 1 answer = {}'.format(part_1(box_ids)))
    print('Part 2 answer = {}'.format(part_2(box_ids)))


if __name__ == '__main__':
    main()
