package book

type BookResponse struct {
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}
