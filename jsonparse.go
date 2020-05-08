package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Users struct
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {

	fmt.Println("Hello World")

	// open json file
	jsonFile, err := os.Open("users.json")

	// if error

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Success reading file")

	defer jsonFile.Close()

	// read opened  jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains
	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		fmt.Println("Twitter Url: " + users.Users[i].Social.Twitter)

	}

}
