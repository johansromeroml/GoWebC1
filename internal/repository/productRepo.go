package repository

import (
	"class1api/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ProductJSONRepo struct {
	products model.ProductList
}

func NewProductJSONRepo() *ProductJSONRepo {
	newRepo := ProductJSONRepo{}
	newRepo.readJSONInfo()
	return &newRepo
}

func (pr *ProductJSONRepo) readJSONInfo() {
	productsFile, err := os.Open("products.json")
	if err != nil {
		println(err.Error())
	}
	defer func(f *os.File) {
		f.Close()
		fmt.Println("File closed")
	}(productsFile)

	productsJSON, err := io.ReadAll(productsFile)
	if err != nil {
		println(err.Error())
	}
	json.Unmarshal(productsJSON, &pr.products.Products)
}

func (r *ProductJSONRepo) writeJSONInfo() {

}

func (r *ProductJSONRepo) productList() (prods model.ProductList, err error) {
	r.readJSONInfo()
	prods = r.products
	return
}

func (r *ProductJSONRepo) productByID(id int) (prod model.Product, err error) {
	r.readJSONInfo()

	return
}

func (r *ProductJSONRepo) saveProduct(newProd model.Product) {
	r.readJSONInfo()
	r.products.Products = append(r.products.Products, newProd)
	r.writeJSONInfo()
}

func (r *ProductJSONRepo) productsPriceGreaterThanValue(value float64) model.ProductList {
	var returnList model.ProductList
	for _, prod := range r.products.Products {
		if prod.Price > value {
			returnList.Products = append(returnList.Products, prod)
		}
	}
	//fmt.Println(len(returnList.Products))
	return returnList
}

func (r *ProductJSONRepo) modifyProduct() {
	// TODO
}
