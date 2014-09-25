# Go AlchemyAPI [![Build Status](https://travis-ci.org/jpadilla/alchemyapi-go.svg?branch=master)](https://travis-ci.org/jpadilla/alchemyapi-go)

Go client library for [AlchemyAPI](http://www.alchemyapi.com/).

## Supported API Calls

- Text Extraction
    - URLGetText
    - URLGetTitle


## Versioning

Each revision of the binding is tagged and the version is updated accordingly.

Given Go's lack of built-in versioning, it is highly recommended you use a
[package management tool](https://code.google.com/p/go-wiki/wiki/PackageManagementTools) in order
to ensure a newer version of the binding does not affect backwards compatibility.

To see the list of past versions, run `git tag`. To manually get an older
version of the client, clone this repo, checkout the specific tag and build the
library:

```sh
git clone https://github.com/jpadilla/alchemyapi-go.git
cd alchemyapi
git checkout api_version_tag
make build
```

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
