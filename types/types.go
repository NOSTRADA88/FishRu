package types

type ProductCard struct {
	Id          int      `json:"id"`
	Price       int32    `json:"price"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Available   bool     `json:"available"`
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
