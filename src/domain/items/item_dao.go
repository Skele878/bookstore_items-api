package items

import (
	"errors"

	"github.com/Skele878/bookstore_items-api/src/clients/elasticsearch"
	"github.com/Skele878/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}
