package ascii

import (
	"fmt"
)

const ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

// Funksjon tar en streng med kun ASCII tegn og lager en utskrift på
// følgende format:
// [ascii-kode heksadesimalt med store bokstaver A-F][mellomrom]
// [symbol for ascii-kode][mellomrom][ascii-kode binært][linjeskift]
// Eksempel (bruk denne funksjonen i en main.go fil):
//	…
// 3E > 111110
// 3F ? 111111
// 40 @ 1000000
// ...
func IterateOverASCIIStringLiteral(sl string) {
	// Kode for Oppgave 1
	//fmt.Println("" + " " + "" + " " + "");

	for i := 0; i < len(sl); i ++ {
		fmt.Print("Hexa = ")
		fmt.Printf("% X\n", sl)
		//fmt.Println()

		fmt.Println("Symbol / input = " + sl)

		fmt.Print("Binært = ")
		fmt.Printf("%b\n", sl[i])
	}
}

// Funksjonen skal generere en utskrift fra en sekvens av bytes,
// dvs. av typen []bytes (det betyr at du må finne den heksadesimale
// eller binære representasjonen av alle tegn i strengen,
// som skal skrives ut (inkludert anførselstegn eller
// “double quotes” på engelsk).
// Funksjonen greetingASCII() returnerer en variabel av typen string,
// som inneholder kun ASCII tegn (ikke utvidet ASCII).
// Gjelder oppgave 1b
func GreetingASCII() {
	//Hello :-) i hexa
	s := "\x48\x65\x6C\x6C\x6F\x20\x3A\x2D\x29"
		fmt.Println(s)
	return s
}