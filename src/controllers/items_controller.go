package controllers

import (
	"fmt"
	"net/http"

	"github.com/Skele878/bookstore_items-api/src/domain/items"
	"github.com/Skele878/bookstore_items-api/src/services"
	"github.com/Skele878/bookstore_oauth-go/oauth"
)

// creating public variable
var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO Return err to the user
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemsService.Create(item)
	if err != nil {
		//TODO return error json to the user
		fmt.Println(result)

		//TODO return created item as json with HTTP status 201 - Created
	}

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {}
