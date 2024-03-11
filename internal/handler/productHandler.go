package handler

import (
	"fmt"
	"net/http"
)

type productHandler struct {
}

func NewProductHandler() *productHandler {
	return &productHandler{}
}

func (ph *productHandler) Pong(wri http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(wri, "pong 2")
}

func (ph *productHandler) ListAllProducts(wri http.ResponseWriter, req *http.Request) {

}
