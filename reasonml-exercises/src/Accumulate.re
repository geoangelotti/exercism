let rec accumulate = (f, x) => switch(x) {
  | [] => []
  | [hd, ...tl] => [f(hd), ...accumulate(f,tl)]
};