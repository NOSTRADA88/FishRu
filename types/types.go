package types

type ProductCard struct {
	Id          int      `json:"id"`
	Price       string   `json:"price"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Available   bool     `json:"available"`
	Category    string   `json:"category"`
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
