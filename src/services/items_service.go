package services

// buisnes logic
import (
	"github.com/Skele878/bookstore_items-api/src/domain/items"
	"github.com/Skele878/bookstore_items-api/src/domain/queries"
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
	Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)
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
func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}

	return dao.Search(query)
}
