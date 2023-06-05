let remainderOf = (year, n) => year mod n == 0;

let isLeapYear = (year) =>
  remainderOf(year, 400) ? { true } : {
    remainderOf(year, 100) ? { false } : {
      remainderOf(year, 4) ? { true } : { false }
	}
};