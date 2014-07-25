package alchemy

type GetTextOptions struct {
	UseMetadata  int
	ExtractLinks int
}

type GetTextResponse struct {
	Status string
	URL    string
	Text   string
}

type GetTitleOptions struct {
	UseMetadata int
}

type GetTitleResponse struct {
	Status string
	URL    string
	Title  string
}
