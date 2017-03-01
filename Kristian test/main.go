package main

import	"./ascii"

func main() {
	var input string;
	//Choose what to translate
	input = "?";
	ascii.IterateOverASCIIStringLiteral(input)
	ascii.GreetingASCII()
}
