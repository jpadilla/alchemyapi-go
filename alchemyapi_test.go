package alchemyapi_test

import (
	"os"
	"testing"

	alchemyapi "github.com/jpadilla/alchemyapi-go"
)

var (
	alchemyAPIKey = os.Getenv("ALCHEMY_API_KEY")
	testURL       = "http://www.cnn.com/2009/CRIME/01/13/missing.pilot/index.html"
)

func init() {
	if len(alchemyAPIKey) == 0 {
		panic("ALCHEMY_API_KEY environment variable is not set, but is needed to run tests!\n")
	}
}

func TestAlchemyAPI_GetTitle(t *testing.T) {
	alchemyClient := alchemyapi.New(alchemyAPIKey)

	titleResponse, err := alchemyClient.GetTitle(testURL, alchemyapi.GetTitleOptions{})

	if err != nil || titleResponse.Status == "ERROR" {
		t.Error(err)
	}
}

func TestAlchemyAPI_GetText(t *testing.T) {
	alchemyClient := alchemyapi.New(alchemyAPIKey)

	textResponse, err := alchemyClient.GetText(testURL, alchemyapi.GetTextOptions{})

	if err != nil || textResponse.Status == "ERROR" {
		t.Error(err)
	}
}
