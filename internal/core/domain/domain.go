package domain

import "time"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
}

type Order struct {
	ID      int        `json:"id"`
	Buyer   string     `json:"buyer"`
	Date    time.Time  `json:"date"`
	Details []ItemLine `json:"details"`
}

type ItemLine struct {
	ID          int    `json:"id"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
