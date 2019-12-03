package main

import (
	"encoding/json"
	"fmt"
)

type bird struct {
	Species     string
	Description string
}

func main() {
	birdJSON := `[
		{
			"Species":"eagle",
			"Description":"bird of prey"
		},
		{
			"Species":"dog",
			"Description":"a dog"
		}
	]`

	var bird []bird
	birdB := []byte(birdJSON)
	json.Unmarshal(birdB, &bird)
	// fmt.Printf("%+v\n", bird)

	b, err := json.Marshal(bird)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}
