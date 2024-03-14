package main

import (
	"fmt"
	"time"

	"github.com/lrtbrabo/go-expert-multithreading/entity"
)

func main() {
  c1 := make(chan entity.BrasilAPI)
  c2 := make(chan entity.ViaCep)
var brasilapi entity.BrasilAPI
var viacep entity.ViaCep

    go brasilapi.MakeRequest("21250500", c1)

    go viacep.MakeRequest("21250500", c2)

  select {
  case msg := <- c1: //BrasilAPI
    fmt.Println("BrasilApi")
    fmt.Println(msg)
  
  case msg := <- c2: //ViaCep
    fmt.Println("ViaCep")
    fmt.Println(msg)

  case <- time.After(time.Second):
    fmt.Println("Timeout")
  }
}
