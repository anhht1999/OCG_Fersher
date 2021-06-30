package model

type Cart struct {
	Id          int     
	CategoryId  int     
	Image       string  
	ProductName string  
	Price       float64 
	Quantity    int     
	CreatedAt   int64   
	ModifiedAt  int64   
}