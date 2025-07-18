package model

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
}

type Subcategory struct {
	ID              int    `json:"id"`
	SubcategoryName string `json:"subcategory_name"`
}
