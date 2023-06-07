let isQuestion = message => message.[String.length(message) - 1] == '?';

let isShouting = message =>
  String.uppercase_ascii(message) == message && Js.Re.test(message, [%re "/[a-z]/i"]);

let hey = (message) =>
  switch (String.trim(message)) {
  | "" => "Fine. Be that way!"
  | message when isQuestion(message) && isShouting(message) => "Calm down, I know what I'm doing!"
  | message when isQuestion(message) => "Sure."
  | message when isShouting(message) => "Whoa, chill out!"
  | _ => "Whatever."
  };