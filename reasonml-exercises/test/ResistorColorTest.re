open Jest;
open Expect;
open ResistorColor;

describe("Resistor Color", () => {
  test("Black", () =>
    expect(colorCode(Black)) |> toEqual(0)
  );

  test("White", () =>
    expect(colorCode(White)) |> toEqual(9)
  );

	test("Orange", () =>
    expect(colorCode(Orange)) |> toEqual(3)
  );
  // Since we have `color` datatype, colors list test from canonical data is skipped
})