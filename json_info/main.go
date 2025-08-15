package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"FullName"`
	Age      int
	Password string `json:"-"`
	Author   string
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// EncodeJson()
	DecodedJson()
}
func EncodeJson() {
	userdata := []User{
		{"Pranav Shinde", 20, "MyPassword", "author", []string{"Mern Stack", "GO Lang developer", "PERN stack"}},
		{"Other User", 23, "His Password", "user", []string{"Java Full Stack", "Python developer", "Mean stack"}}}

	// package this data as json data
	json_data, err := json.MarshalIndent(userdata, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json_data)

}

func DecodedJson() {
	jsonData := []byte(`
	{
	"FullName": "Pranav Shinde",
	"Age": 20,
	 "Password":"MyPassword",
	 "Author":"admin"
	 }
		`)

	var data User

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("Json is Valid")
		json.Unmarshal(jsonData, &data)
		fmt.Printf("%#v\n", data)
	} else {
		fmt.Println("Json is not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v", myOnlineData)
}
