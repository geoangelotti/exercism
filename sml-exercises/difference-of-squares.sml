fun square x = x * x;

fun add (x,y) = x + y;

fun create n = List.tabulate (n, fn x => x + 1);

fun squareOfSum n = 
    square (List.foldl ( add ) 0 (create n));

fun sumOfSquares n =
    List.foldl ( add ) 0 (List.map square (create n));


fun differenceOfSquares n =
    (squareOfSum n) - (sumOfSquares  n);
