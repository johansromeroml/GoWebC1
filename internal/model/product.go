package model

type Product struct {
	ID             int     `json:"id" default:"-1"`
	Name           string  `json:"name"`
	Quantity       int     `json:"quantity"`
	CodeValue      string  `json:"code_value"`
	Published      bool    `json:"is_published" default:"false"`
	ExpirationDate string  `json:"expiration"`
	Price          float64 `json:"price"`
}

type ProductList struct {
	Products []Product
}

func (pl *ProductList) AddProduct(prod Product) {
	pl.Products = append(pl.Products, prod)
}

func (pl *ProductList) getProductByID(id int) (prod Product) {
	for _, product := range pl.Products {
		if prod.ID == id {
			prod = product
		}
	}
	return
}
