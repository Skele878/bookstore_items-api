package services

// buisnes logic
import (
	"net/http"

	"github.com/Skele878/bookstore_items-api/src/domain/items"
	"github.com/Skele878/bookstore_utils-go/rest_errors"
)

// Creating new variable for exporting and using methods of interface
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

// implementing methods of the inteface
func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}
func (s *itemsService) Get(string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "not_implemented", nil)
}
