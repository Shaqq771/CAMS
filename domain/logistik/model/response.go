package model

type AddedProductResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DeletedProductResponse struct {
	Id int `json:"id"`
}

type ProductLists struct {
	Pagination Pagination `json:"pagination"`
	Product    []Product  `json:"products"`
	Sort       []string   `json:"sort,omitempty"`
	Filter     []string   `json:"filter,omitempty"`
}
