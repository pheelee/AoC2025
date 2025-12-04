import os

# input file handling
script_dir = os.path.dirname(os.path.abspath(__file__))
input_file = os.path.join(script_dir, "input-python.txt")

# initialize total invalid ID sum
inv_id_sum = 0

# read input file and split by comma to get individual ranges
with open(input_file, "r") as input_file:
    input_data = (input_file.read().strip().splitlines())[0].split(",")

# process each range    
for r in input_data:
    # print(f"Range: {r}")
    # get range start and end
    r_start, r_end = map(int, r.split("-"))

    # iterate through all the numbers within the range
    for id in range(r_start, r_end + 1):
        # print(f"ID: {id}, digits: {len(str(id))}")
        # get the amount of digits in the number
        digits = len(str(id))

        # if digits are even process further, else skip to next number
        if digits % 2 == 0:
            # split the number in two halves
            first_half = int(str(id)[:digits//2])
            second_half = int(str(id)[digits//2:])

            # print(f"  first half: {first_half}, second half: {second_half}")

            # if first half equals second half, we found an invalid ID and add them up
            if first_half == second_half:
                # print(f"  invalid ID found: {id}")
                inv_id_sum += id

print(f"Part I: {inv_id_sum}")
 


