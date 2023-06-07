fun accumulate (f, x::xs) = f(x)::accumulate(f,xs)
    | accumulate (f, nil) = []
