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
	birdJSON := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`

	var bird []bird
	json.Unmarshal([]byte(birdJSON), &bird)
	fmt.Printf("%+v\n", bird)
}
