import os

# input file handling
script_dir = os.path.dirname(os.path.abspath(__file__))
input_file = os.path.join(script_dir, "input-python.txt")

# initialize dial range start, end and calculcate size
r_start = 0
r_end = 99
r_size = r_end - r_start + 1

dial_value = 50
zero_counter = 0
rotations = 0

with open(input_file, "r") as input_file:    
    for line in input_file:
        last_dial_value = dial_value
        if line[0] == "R":
            dial_value = dial_value + int(line[1:])
        elif line[0] == "L":
            dial_value = dial_value - int(line[1:])
        
        if dial_value > r_end or dial_value <= r_start:
            # if we turn right, divide dial value (befor mod) by size of range
            if line[0] == "R":
                rotations += int(dial_value / r_size)

            # if we turn left, add range size to the absolute dial value before dividing (to basically set)
            elif line[0] == "L":
                rotations += int((abs(dial_value) + r_size) / r_size)

                # if we turn left and dial was on 0 in the last turn, we need to subtract 1 rotation
                if last_dial_value == 0:
                    rotations -= 1

            dial_value = dial_value % r_size
        
        if dial_value == 0:
            zero_counter += 1

print(f"Part I: {zero_counter}")
print(f"Part II: {rotations}")
