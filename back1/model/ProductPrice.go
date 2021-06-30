package model

type ProductPrice struct {
	ID uint
	ProductID uint
	Product *Product
	Price float32
	Discount int
	CreatedAt int
}