package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	// "github.com/lrtbrabo/go-expert-multithreading/entity"
)

func MakeRequestBrasilAPI(cep string, ch chan string) {
	uri := "https://brasilapi.com.br/api/cep/v1/" + cep

	c := http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	ch <- string(body)
}

func MakeRequestViaCep(cep string, ch chan string) {
	uri := "http://viacep.com.br/ws/" + cep + "/json/"

	c := http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	ch <- string(body)
}

func main() {
  // c_cep := make(chan string)
  c1 := make(chan string)
	c2 := make(chan string)
  cep := "21020122"

	go MakeRequestBrasilAPI(cep, c1)
	go MakeRequestViaCep(cep, c2)

	select {
	case msg := <-c1: //BrasilAPI
		fmt.Println(msg)
		fmt.Println("Enviado por: BrasilAPI")

	case msg := <-c2: //ViaCep
		fmt.Println(msg)
		fmt.Println("Enviado por: ViaCep")

	case <-time.After(time.Second):
		fmt.Println(errors.New("Operation timed out"))
	}
}

// func main() {
// 	c1 := make(chan entity.BrasilAPI)
// 	c2 := make(chan entity.ViaCep)
// 	var brasilapi entity.BrasilAPI
// 	var viacep entity.ViaCep
//
// 	go brasilapi.MakeRequest("21250500", c1)
// 	go viacep.MakeRequest("21250500", c2)
//
// 	select {
// 	case msg := <-c1: //BrasilAPI
// 		fmt.Println("BrasilApi")
// 		fmt.Println(msg)
//
// 	case msg := <-c2: //ViaCep
// 		fmt.Println("ViaCep")
// 		fmt.Println(msg)
//
// 	case <-time.After(time.Second):
// 		fmt.Println("Timeout")
// 	}
// }
