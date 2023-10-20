package types

type ProductCard struct {
	Id          int      `json:"id"`
	Price       string   `json:"price"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Category    string   `json:"category"`
	Slug        string   `json:"slug"`
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
