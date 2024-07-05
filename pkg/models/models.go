package models

type Request struct {
	Method  int
	URL     string
	Headers []byte
}

type Response struct {
	ID      int
	Status  int
	Headers []byte
	Length  int
}
