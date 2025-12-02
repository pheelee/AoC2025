import os

# input file handling
script_dir = os.path.dirname(os.path.abspath(__file__))
input_file = os.path.join(script_dir, "input-python.txt")

# print(input_file)

r_start = 0
r_end = 99
r_size = r_end - r_start + 1

dial_value = 50
zero_counter = 0

with open(input_file, "r") as input_file:
    for line in input_file:
        if line[0] == "R":
            dial_value = dial_value + int(line[1:])
        elif line[0] == "L":
            dial_value = dial_value - int(line[1:])
        
        if dial_value > r_end or dial_value < r_start:
            dial_value = dial_value % r_size
        
        if dial_value == 0:
            zero_counter += 1

print(f"Part I: {zero_counter}")

# Part II: coming soon...
