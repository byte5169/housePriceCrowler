package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// We create a file, marshal the data into it, and write it to a file
//
// Args:
//   data ([]House): The data to be written to the file.
func WriteJSON(data []House) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}
	_ = ioutil.WriteFile("houses.json", file, 0644)
}
