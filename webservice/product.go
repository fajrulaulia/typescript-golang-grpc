package webservice

import (
	"context"
	"errors"
	"webservice/product"
)

type API struct {
	product.UnimplementedProductServiceServer
}

func (s *API) ProductHandler(ctx context.Context, req *product.ProductRequest) (*product.ProductResponse, error) {
	return Catalogue(req.Id)
}

func Catalogue(id string) (*product.ProductResponse, error) {
	var res product.ProductResponse
	switch id {
	case "1":
		res.Id = "1"
		res.Name = "baju hari raya"
		res.Price = 34000
	case "2":
		res.Id = "2"
		res.Name = "baju anak-anak"
		res.Price = 60000

	case "3":
		res.Id = "3"
		res.Name = "seragam sekolah"
		res.Price = 90000

	case "4":
		res.Id = "4"
		res.Name = "baju hari keagamaan"
		res.Price = 122000

	case "5":
		res.Id = "5"
		res.Name = "seragam dinas"
		res.Price = 923000

	default:
		return nil, errors.New("error not found")
	}

	return &res, nil
}
