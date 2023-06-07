// Generated by ReScript, PLEASE EDIT WITH CARE
'use strict';

var Jest = require("@glennsl/bs-jest/src/jest.bs.js");
var ResistorColor = require("../src/ResistorColor.bs.js");

Jest.describe("Resistor Color", (function (param) {
        Jest.test("Black", (function (param) {
                return Jest.Expect.toEqual(0, Jest.Expect.expect(ResistorColor.colorCode(/* Black */0)));
              }));
        Jest.test("White", (function (param) {
                return Jest.Expect.toEqual(9, Jest.Expect.expect(ResistorColor.colorCode(/* White */9)));
              }));
        return Jest.test("Orange", (function (param) {
                      return Jest.Expect.toEqual(3, Jest.Expect.expect(ResistorColor.colorCode(/* Orange */3)));
                    }));
      }));

/*  Not a pure module */