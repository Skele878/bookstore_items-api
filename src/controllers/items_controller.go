package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Skele878/bookstore_items-api/src/domain/items"
	"github.com/Skele878/bookstore_items-api/src/services"
	httputils "github.com/Skele878/bookstore_items-api/src/utils/http_utils"
	"github.com/Skele878/bookstore_oauth-go/oauth"
	"github.com/Skele878/bookstore_utils-go/rest_errors"
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

// autentification request
func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	//if we have err
	if err := oauth.AuthenticateRequest(r); err != nil {
		httputils.RespondError(w, err)
		return
	}
	// if we dont have accesstoken yet we are not able to use CallerId to validate userID

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("invalid access token")
		httputils.RespondError(w, respErr)
		return
	}

	// if the accestoken valid and request = Ok
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		httputils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	//use requestBody to fill  the itemRequest
	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json body")
		httputils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = oauth.GetCallerId(r)

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		httputils.RespondError(w, createErr)
		return
	}

	httputils.RespondJson(w, http.StatusCreated, result)

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {}
