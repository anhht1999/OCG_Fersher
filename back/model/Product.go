package model

type Product struct {
	ID          uint 
	Name        string
	Description string
	Author      string
	Publisher   string
	Supplier    string
	Year        uint
	Page        uint
	Price       float64
	Url       	string
	CategoryID  uint
	Category    *Category
}

//`ID`, `name`, `description`, `author`, `publisher`, `supplier`, `year`, `page`, `price`, `category_id`
