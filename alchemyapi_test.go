package alchemyapi_test

import (
	"log"

	alchemyapi "github.com/jpadilla/alchemyapi-go"
)

func ExampleAlchemyAPI_GetTitle() {
	alchemyAPIKey = "ALCHEMY_API_KEY"
	alchemyClient := alchemyapi.New(alchemyAPIKey)

	titleResponse, err := alchemyClient.GetTitle(data.URL, alchemyapi.GetTitleOptions{})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", titleResponse.Title)
}

func ExampleAlchemyAPI_GetText() {
	alchemyAPIKey = "ALCHEMY_API_KEY"
	alchemyClient := alchemyapi.New(alchemyAPIKey)

	textResponse, err := alchemyClient.GetText(data.URL, alchemyapi.GetTextOptions{})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", textResponse.Text)
}
