package database

type ProductCard struct {
	Id          int      `json:"id"`
	Price       int32    `json:"price"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Available   bool     `json:"available"`
}
