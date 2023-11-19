package main

import (
	data "cs-market-db.danwebr.net/internal"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (app *application) showSteamPriceOverview(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Country        string `json:"country"`
		Currency       string `json:"currency" validate:"required"`
		Appid          string `json:"appid" validate:"required"`
		MarketHashName string `json:"market_hash_name" validate:"required"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	url := "https://steamcommunity.com/market/priceoverview/?country" + input.Country + "&currency=" + input.Currency + "&appid=" + input.Appid + "&market_hash_name=" + input.MarketHashName

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var priceOverview data.PriceOverview
	err = json.Unmarshal(body, &priceOverview)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"priceOverview": priceOverview}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
