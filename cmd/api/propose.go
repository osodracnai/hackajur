package main

import (
	"github.com/osodracnai/hackajur/internal/database"
	"github.com/osodracnai/hackajur/internal/request"
	"net/http"
)

func (app *application) createPropose(w http.ResponseWriter, r *http.Request) {
	input := database.Propose{}
	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	w.Write([]byte("This is a protected handler"))
}
