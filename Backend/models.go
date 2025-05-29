package main

type Product struct {
	ID       int     `json:"id"`
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
