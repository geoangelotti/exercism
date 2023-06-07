// Generated by ReScript, PLEASE EDIT WITH CARE
'use strict';

var $$String = require("bs-platform/lib/js/string.js");
var Caml_string = require("bs-platform/lib/js/caml_string.js");

function isQuestion(message) {
  return Caml_string.get(message, message.length - 1 | 0) === /* '?' */63;
}

function isShouting(message) {
  if ($$String.uppercase_ascii(message) === message) {
    return /[a-z]/i.test(message);
  } else {
    return false;
  }
}

function hey(message) {
  var message$1 = $$String.trim(message);
  if (message$1 === "") {
    return "Fine. Be that way!";
  } else if (isQuestion(message$1) && isShouting(message$1)) {
    return "Calm down, I know what I'm doing!";
  } else if (isQuestion(message$1)) {
    return "Sure.";
  } else if (isShouting(message$1)) {
    return "Whoa, chill out!";
  } else {
    return "Whatever.";
  }
}

exports.hey = hey;
/* No side effect */