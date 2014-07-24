package alchemy

type Options struct {
	UseMetadata  int
	ExtractLinks int
}

type Response struct {
	Status string
	URL    string
	Text   string
}
