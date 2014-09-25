# Go AlchemyAPI

Go client library for [AlchemyAPI](http://www.alchemyapi.com/).

## Supported API Calls

- Text Extraction
    - URLGetText
    - URLGetTitle


## Installation

```
go get github.com/jpadilla/alchemyapi-go
```

## Documentation

For details on all the functionality in this library, see the [GoDoc](http://godoc.org/github.com/jpadilla/alchemyapi-go) documentation.

## Example usage

```go
package main

import (
    "fmt"
    "log"

    alchemyapi "github.com/jpadilla/alchemyapi-go"
)

func main() {
    alchemyAPIKey = "ALCHEMY_API_KEY"
    alchemyClient := alchemyapi.New(alchemyAPIKey)

    titleResponse, err := alchemyClient.GetTitle(data.URL, alchemyapi.GetTitleOptions{})

    if err != nil {
        log.Fatal(err)
    }

    log.Printf("%v\n", titleResponse.Title)

    textResponse, err := alchemyClient.GetText(data.URL, alchemyapi.GetTextOptions{})

    if err != nil {
        log.Fatal(err)
    }

    log.Printf("%v\n", textResponse.Text)
}

```
