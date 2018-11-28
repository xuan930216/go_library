package files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func ParseFile() {
	//read file using the os package.
	//once we have opened the file, we will defer the close until the end of the function
	pwd, _ := os.Getwd()
	jsonFile, err := os.Open(pwd + "/utils/files/users.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened the file")
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var users Users
	json.Unmarshal(byteValue, &users)

	for _, item := range users.Users {
		fmt.Println("User Type: ", item.Type)
		fmt.Println("User age: ", item.Age)
		fmt.Println("User name: ", item.Name)
		fmt.Println("Facebook Url: ", item.Social.Facebook)
	}
}
