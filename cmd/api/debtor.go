package main

import (
	"github.com/osodracnai/hackajur/internal/request"
	"github.com/osodracnai/hackajur/internal/response"
	"net/http"
)

type Address struct {
	PostalCode string `json:"postalCode"`
	City       string `json:"city"`
	Uf         string `json:"uf"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
}
type Debtor struct {
	ID             string  `json:"id"`
	FiscalDocument string  `json:"fiscalDocument"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	TypeOfDebtor   string  `json:"typeOfDebtor"`
	Address        Address `json:"address"`
}

func (app *application) createDebtor(w http.ResponseWriter, r *http.Request) {
	input := Debtor{}
	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	id, err := app.db.InsertDebtor(input.FiscalDocument, input.Name, input.Email,
		input.TypeOfDebtor, input.Address.PostalCode, input.Address.City,
		input.Address.Uf, input.Address.Number, input.Address.Complement)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, id)
	if err != nil {
		app.serverError(w, r, err)
	}
}
