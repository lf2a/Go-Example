package main

import (
	"encoding/json"
	"fmt"
)

type Notebook struct {
	Marca  string
	Modelo string
	Preco  string
}

func main() {

	notebooksString := `[
		{
			"marca":"Acer",
			"modelo":"Aspire 5",
			"preco":"1234.90"
		},
		{
			"marca":"Acer",
			"modelo":"Aspire Nitro 5",
			"preco":"3234.90"
		},
		{
			"marca":"Apple",
			"modelo":"Macbook air",
			"preco":"11234.90"
		}
	]`
	var notebooks []Notebook

	jsonString := []byte(notebooksString)

	// convertendo json string em struct
	json.Unmarshal(jsonString, &notebooks)
	fmt.Println(notebooks)

	// convertendo struct em json
	stringJson, err := json.Marshal(notebooks)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(stringJson))

	fmt.Printf("%+v\n", notebooks)
}
