package domain

type Stock struct {
	Symbol string  `csv:"Symbol"`
	Date   string  `csv:"Date"`
	Time   string  `csv:"Time"`
	Open   float32 `csv:"Open"`
	High   float32 `csv:"High"`
	Low    float32 `csv:"Low"`
	Close  float32 `csv:"Close"`
	Volume int     `csv:"Volume"`
	Name   string  `csv:"Name"`
}
