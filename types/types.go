package types

type ProductCard struct {
	Id           int      `json:"id"`
	Price        string   `json:"price"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Photos       []string `json:"photos"`
	Category     string   `json:"category"`
	SlugName     string   `json:"slugName"`
	SlugCategory string   `json:"slugCategory"`
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
