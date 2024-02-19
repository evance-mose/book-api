package datatypes

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Book struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Books struct {
	Books []Book `json:"books"`
}

func (bk *Books) AddBook(book Book) {
	bk.Books = append(bk.Books, book)
}
