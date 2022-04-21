package utils

// It takes a struct of type House and returns a string
//
// Args:
//   houseInfo (House): This is the struct that contains all the information about the house.
//
// Returns:
//   A string
import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

// Link to otodom.pl with preset params for city Poznan
var webLink = "https://www.otodom.pl/pl/oferty/wynajem/mieszkanie/poznan?page=1&limit=%v&market=ALL&ownerTypeSingleSelect=ALL"

type House struct {
	Name                  string
	Description           string
	MapLink               string
	RentLink              string
	Price                 string
	Image                 string
	HouseDescriptionTable map[string]string
}

// It visits a website, collects all the links to the houses, visits each of them, collects the data
// and returns it
//
// Returns:
//   A slice of structs
func Crawl() []House {

	allHouses := make([]House, 0)
	allHousesUrls := []string{}

	// number of latest houses to check
	housesToCheck := "10"

	// create Collectors
	// async can be used, but didnt figure out how to merge crowler and saving data to
	// slice of structs so that data isnt mixed up
	c := colly.NewCollector(colly.AllowedDomains("otodom.pl", "www.otodom.pl"))
	c.Limit(&colly.LimitRule{
		RandomDelay: 10 * time.Second,
	})
	infoCollector := c.Clone()
	infoCollector.Limit(&colly.LimitRule{
		RandomDelay: 10 * time.Second,
	})

	c.OnHTML("ul.css-14cy79a.e3x1uf06", func(e *colly.HTMLElement) {

		e.ForEach("li.css-p74l73", func(_ int, h *colly.HTMLElement) {
			houseUrl := h.ChildAttr("a", "href")
			url := h.Request.AbsoluteURL(houseUrl)
			allHousesUrls = append(allHousesUrls, url)
		})

		for index, url := range allHousesUrls[3:] {
			fmt.Print("Visiting ", index+1, ": ", url, "\n\n\n")
			infoCollector.Visit(url)
		}
	})

	// create temp struct to store data
	tmpHouse := House{}
	infoCollector.OnRequest(func(request *colly.Request) {
		tmpHouse.MapLink = request.URL.String() + "#map"
		tmpHouse.RentLink = request.URL.String()

	})

	infoCollector.OnHTML("div.e1t9fvcw3", func(e *colly.HTMLElement) {

		houseDescriptionTable := make(map[string]string)
		e.ForEach(".estckra9", func(_ int, h *colly.HTMLElement) {
			keyElement := h.ChildText("div.estckra7")
			valueElement := h.ChildText("div.estckra8 > div.estckra5")
			houseDescriptionTable[keyElement] = valueElement
		})

		tmpHouse.Name = e.ChildText("h1.css-11kn46p")
		tmpHouse.Price = e.ChildText("strong.css-8qi9av")
		tmpHouse.Description = e.ChildText("div.e1r1048u1 > div")
		tmpHouse.Image = e.ChildAttr("img", "src")
		tmpHouse.HouseDescriptionTable = houseDescriptionTable

		// append all data to slice of structs
		allHouses = append(allHouses, tmpHouse)

	})

	link := fmt.Sprintf(webLink, housesToCheck)
	c.Visit(link)

	return allHouses
}
