BEGIN {
    FS = ","    # comma is the field separator
}

# Task 1: When the first field is "needs_license",
#         print "true" if the second field contains "car" or "truck".

$1 ~ /needs_license/ && $2 ~ /car|truck/ { print "true" }

# Task 2: When the first field is "resell_price",
#         print the resell value of the original price (second field)
#         given the age of the car (third field).

$1 ~ /resell_price/ && $3 < 3 && $3 >= 0  { print $2 * 0.8 }
$1 ~ /resell_price/ && $3 <= 10 && $3 >= 3 { print $2 * 0.7 }
$1 ~ /resell_price/ && $3 > 10 { print $2 * 0.5 }
