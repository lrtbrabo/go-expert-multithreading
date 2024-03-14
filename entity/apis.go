package entity

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type BrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type Api interface{
  GetReponse()
  MakeRequest()
}

func (b *BrasilAPI) GetResponse(response string) BrasilAPI{
  var result BrasilAPI
  err := json.Unmarshal([]byte(response), &result)
  if err != nil {
    log.Fatal(err)
  }
  return result
}

func (v *ViaCep) GetResponse(response string) ViaCep{
  var result ViaCep
  err := json.Unmarshal([]byte(response), &result)
  if err != nil {
    log.Fatal(err)
  }
  return result
}


func (b *BrasilAPI) MakeRequest(cep string, ch chan BrasilAPI) {
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
  
  ch <- b.GetResponse(string(body))
}

func (v *ViaCep) MakeRequest(cep string, ch chan ViaCep) {
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
  ch <- v.GetResponse(string(body))
}


