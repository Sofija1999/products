package models

import "time"

type Product struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Quantity int64 `json:"quantity"`
}