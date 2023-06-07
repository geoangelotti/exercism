let divi = (n: int, i: int, msg: string) => n mod i == 0 ? Some(msg) : None

let clean = fun 
	| Some(_) => true
	| None => false

let clean2 = fun 
	| Some(message) => message
	| None => ""

let raindrops = n => {
  let messagesOption = [divi(n, 3, "Pling"), divi(n, 5, "Plang"), divi(n, 7, "Plong")];
  let messages = List.filter(clean, messagesOption);
  List.length(messages) > 0 ? List.fold_left((++), "", List.map(clean2, messagesOption)) : string_of_int(n)
}