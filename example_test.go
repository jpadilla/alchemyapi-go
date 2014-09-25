package alchemyapi_test

import (
	"log"

	alchemyapi "github.com/jpadilla/alchemyapi-go"
)

func ExampleAlchemyAPI_GetTitle() {
	alchemyAPIKey := "ALCHEMY_API_KEY"
	alchemyClient := alchemyapi.New(alchemyAPIKey)
	url := "http://www.cnn.com/2009/CRIME/01/13/missing.pilot/index.html"

	titleResponse, err := alchemyClient.GetTitle(url, alchemyapi.GetTitleOptions{})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", titleResponse.Title)
}

func ExampleAlchemyAPI_GetText() {
	alchemyAPIKey := "ALCHEMY_API_KEY"
	alchemyClient := alchemyapi.New(alchemyAPIKey)
	url := "http://www.cnn.com/2009/CRIME/01/13/missing.pilot/index.html"

	textResponse, err := alchemyClient.GetText(url, alchemyapi.GetTextOptions{})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", textResponse.Text)
}
