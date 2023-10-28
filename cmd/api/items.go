package main

import (
	data "cs-market-db.danwebr.net/internal"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createItemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new item")
}

func (app *application) showItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	item := data.Items{
		ID:         id,
		CreatedAt:  time.Now(),
		Name:       "Karambit - Fade",
		Case:       0,
		Rarity:     5,
		Conditions: []string{"Factory New", "Minimal Wear"},
		Version:    1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"item": item}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
