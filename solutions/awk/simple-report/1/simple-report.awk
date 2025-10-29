BEGIN {
    FS = ","
}

{
    firstNumber = $3 $4
    secondNumber = $5 $6
    average = (firstNumber + secondNumber) / 2
    print "#"$1", "$2" = " average
}
