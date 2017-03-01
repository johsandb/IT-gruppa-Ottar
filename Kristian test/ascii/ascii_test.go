package ascii

import "testing"

func TestGreetingASCII(t *testing.T) {
	if !(isASCII("æ")){		//Expected: Fail
	//if !(isASCII("S")){		//Expected: Ok / Success
		t.Fail()
	}
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}

//Implementer en test for funksjonen greetingASCII() i egen fil
//ascii_test.go, som tester om returverdier (av type string) inneholder
//kun ASCII-tegn.
