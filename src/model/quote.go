package model

type QuoteData struct {
	Id         string   `json:"id"`
	Quote      string   `json:"quote"`
	Length     string   `json:"length"`
	Author     string   `json:"author"`
	Tags       []string `json:"tags"`
	Category   string   `json:"category"`
	Date       string   `json:"date"`
	Permalink  string   `json:"parmalink"`
	Title      string   `json:"title"`
	Background string   `json:"Background"`
}

type QuoteResponse struct {
	Contents Contents `json:"contents"`
}

type Contents struct {
	Quotes []Quote `json:"quotes"`
}

type Quote struct {
	Quote string `json:"quote"`
}

type APISuccess struct {
	Total string `json:"total"`
}
