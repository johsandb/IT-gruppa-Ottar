package main

import (
  "fmt"
  "os"
  "strconv"
)

func main(){
  //Her har vi koden fra oppgave linken, men
  //vi får denne feilmeldingen
  //strconv.ParseFloat: parsing "E:\\Christoffer\\Documents\\bin\\Test.exe": invalid syntax
  //Args[1] setter antall parametere
  //Så bruker vi ParseFloat for å kovertere inputen som er String
  //til int ellr tall som da kan brukes til regning i lnije 24
    arg1 := os.Args[1]
		arg2 := os.Args[2]
    argInput1, err := strconv.ParseFloat(arg1, 64)
		argInput2, err := strconv.ParseFloat(arg2, 64)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Dette er svaret dit:")
    fmt.Println(argInput1 + argInput2)
}
