package types

type Login struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"author"`
	Role     string `json:"role"`
	Success  int8   `json:"success"`
}

type Book struct {
	BookId   uint   `json:"bookid"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Shelf    string `json:"shelf"`
	IssuedBy int64  `json:"issuedby"`
	Success  int8   `json:"success"`
}

type Data struct {
	Page   string
	Result any
}
