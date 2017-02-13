package main

import (
  "fmt"
  "os"
  "strconv"
)

func main(){
    arg1 := os.Args[1]
		arg2 := os.Args[2]
    argInput1, err := strconv.ParseFloat(arg1, 64)
		argInput2, err := strconv.ParseFloat(arg2, 64)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Summen av " + argInput1 + " + " + argInput2 + " er:")
    fmt.Println(argInput1 + argInput2)
}
