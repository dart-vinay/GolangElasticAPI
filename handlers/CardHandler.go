package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "../models"
	helper "../util"
	"github.com/gorilla/mux"
	"github.com/olivere/elastic"
)

func GetAllCardsForUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside GetAllCardsForUser method")

	vars := mux.Vars(r)
	userID := vars["userID"]
	var cards []model.Card

	numberOfRecords := 50

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMultiMatchQuery(userID, "createdBy"))

	/// Validate Query
	isValid := helper.ValidateQuery(searchSource)

	if !isValid {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	searchService := ESClientGlobal.Search().Index("cards").SearchSource(searchSource).Size(numberOfRecords)
	searchResult, err := searchService.Do(ctxGlobal)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Invalid Result", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var card model.Card
		err := json.Unmarshal(hit.Source, &card)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error unmarshalling", err)
		}
		cards = append(cards, card)
	}

	if err != nil {
		fmt.Println("Fetching Cards fail: ", err)
	} else {
		for _, s := range cards {
			fmt.Printf("Cards found cardID: %s, Latitude: %f, Longitude: %f \n", s.CardId, s.Longitude, s.Latitude)
		}
	}
	fmt.Println("Success!")
	json.NewEncoder(w).Encode(cards)

}

func GetTopCards(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside GetTopCards method")

	vars := mux.Vars(r)
	var cards []model.Card
	var result []model.Card
	numberOfRecords := 200

	//Validate Input
	latitude, longitude, validParams := helper.ValidateInput(vars["latitude"], vars["longitude"])

	if !validParams {
		fmt.Println("Invalid Params")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	searchSource := elastic.NewSearchSource().Query(elastic.NewMatchAllQuery())

	// Validate Query
	isValid := helper.ValidateQuery(searchSource)

	if !isValid {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	searchService := ESClientGlobal.Search().Index("cards").SearchSource(searchSource).Size(numberOfRecords)
	searchResult, err := searchService.Do(ctxGlobal)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Invalid Result", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var card model.Card
		err := json.Unmarshal(hit.Source, &card)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error unmarshalling", err)
		}
		cards = append(cards, card)
	}

	if err != nil {
		fmt.Println("Fetching cards fail: ", err)
	} else {
		fmt.Println("Success!")
	}

	helper.GetTopTenCards(&result, cards, longitude, latitude)

	json.NewEncoder(w).Encode(result)
}
