package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Write a program which prompts the user to first enter a name,
// and then enter an address. Your program should create
//  a map and add the name and address to the map using the keys
//  “name” and “address”, respectively. Your program should use
//  Marshal() to create a JSON object from the map, and then your
//  program should print the JSON object.

func main() {
	jsonObject := map[string]string{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your name.")
	name, _, _ := reader.ReadLine()
	jsonObject["name"] = string(name)

	fmt.Println("Please enter your address.")
	address, _ := reader.ReadString('\n')
	jsonObject["address"] = strings.Trim(address, "\n")

	barr, _ := json.Marshal(jsonObject)

	fmt.Println(string(barr))

}
