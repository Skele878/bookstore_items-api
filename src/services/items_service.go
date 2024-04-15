package services

import (
	"net/http"

	"github.com/Skele878/bookstore_items-api/src/domain/items"
	"github.com/Skele878/bookstore_utils-go/rest_errors"
)

// Creating new variable
var (
	ItemsService itemsServiceInterface = &itemsService{}
)

// basic interface for service
type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct {
}

// imlementing methods of the inteface
func (s *itemsService) Create(items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "not_implemented", nil)
}
func (s *itemsService) Get(string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "not_implemented", nil)
}
