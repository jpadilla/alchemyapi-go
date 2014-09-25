package alchemyapi

// GetTextOptions is the set of parameters that can be used on the URLGetText endpoint.
// For more details see http://www.alchemyapi.com/api/text/urls.html#text.
type GetTextOptions struct {
	UseMetadata  int
	ExtractLinks int
}

// GetTextResponse is the resource representing response from URLGetText endpoint.
// For more details see http://www.alchemyapi.com/api/text/urls.html#text.
type GetTextResponse struct {
	Status string
	URL    string
	Text   string
}

// GetTitleOptions is the set of parameters that can be used on the URLGetTitle endpoint.
// For more details see http://www.alchemyapi.com/api/text/urls.html#title.
type GetTitleOptions struct {
	UseMetadata int
}

// GetTitleResponse is the resource representing response from URLGetTitle endpoint.
// For more details see http://www.alchemyapi.com/api/text/urls.html#title.
type GetTitleResponse struct {
	Status string
	URL    string
	Title  string
}
