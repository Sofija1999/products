package models

import "time"

type Product struct {
	id int64 `json:"id"`
	price float64 `json:"price"`
	shortDescription string `json:"shortDescription"`
	description string `json:"description"`
	created time.Time `json:"created"`
	updated time.Time `json:"updated"`
	name string `json:"name"`
}