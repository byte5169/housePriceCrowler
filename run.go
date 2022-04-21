package main

import (
	"fmt"
	"otodom/utils"
)

// The function crawls the website, writes the data to a JSON file, gets the names of the houses from
// the Firestore database, checks if the house is already in the database, if not, it forms a message
// and sends it to the user
func main() {

	// get latest data
	data := utils.Crawl()

	// write it to JSON file
	utils.WriteJSON(data)

	// get list of all houses in DB
	// may be it can be done pretier but didnt have much time
	housesNames := utils.GetDocsNameFirestore()

	// for each house check if it is new in DB or not,
	// if not => send to telegram bot
	for index := range data {

		utils.WriteFirestore(data[index])
		checkResult := utils.Contains(housesNames, data[index].Name)

		if checkResult {
			fmt.Println("Nothing new")
		} else {
			message := utils.FormMessage(data[index])
			utils.SendPM(message)
		}
	}

}
