package utils

// It takes a struct of type House and returns a string
//
// Args:
//   houseInfo (House): This is the struct that contains all the information about the house.
//
// Returns:
//   A string
func FormMessage(houseInfo House) string {

	message := "Name: " + houseInfo.Name + "\n" + "Price: " + houseInfo.Price + "\n" + "Map Link: " + houseInfo.MapLink + "\n" + "House Link: " + houseInfo.RentLink + "\n"
	return message
}
