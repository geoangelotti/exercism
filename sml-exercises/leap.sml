fun remainder_of year n = (year mod n = 0)

fun isLeapYear year =
    if remainder_of year 400 then true else if remainder_of year 100 then false else if remainder_of year 4 then true else false
