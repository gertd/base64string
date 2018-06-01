package main

import (
	"encoding/json"
	"fmt"
	"log"

	b64 "github.com/gertd/base64string"
)

type user struct {
	ID    string           `json:"id"`
	Name  b64.Base64String `json:"name"`
	Email b64.Base64String `json:"email"`
}

func main() {

	u := user{
		ID:    "965DE03F-D535-4F26-A1B6-2AAF23BF75BF",
		Name:  b64.Base64String("Milo Hoffman"),
		Email: b64.Base64String("milo.hoffman@antitrust.movie"),
	}

	b, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	var u2 user
	json.Unmarshal(b, &u2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("raw:   %v\n\n", u2)

	fmt.Printf("name:  %s\n", u2.Name.Get())
	fmt.Printf("email: %s\n", u2.Email.Get())
}
