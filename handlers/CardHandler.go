package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	// To verify the query
	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}
	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))

	searchService := ESClientGlobal.Search().Index("cards").SearchSource(searchSource).Size(numberOfRecords)
	searchResult, err := searchService.Do(ctxGlobal)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var card model.Card
		err := json.Unmarshal(hit.Source, &card)
		if err != nil {
			fmt.Println("[Getting Cards][Unmarshal] Err=", err)
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

	json.NewEncoder(w).Encode(cards)

}

func GetTopCards(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside GetTopCards method")
	vars := mux.Vars(r)

	var cards []model.Card
	var result []model.Card
	numberOfRecords := 200

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchAllQuery())

	// To verify the query
	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}
	fmt.Println("[esclient]Final ESQuery=", string(queryJs))

	searchService := ESClientGlobal.Search().Index("cards").SearchSource(searchSource).Size(numberOfRecords)
	searchResult, err := searchService.Do(ctxGlobal)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var card model.Card
		err := json.Unmarshal(hit.Source, &card)
		if err != nil {
			fmt.Println("[Getting Cards][Unmarshal] Err=", err)
		}
		cards = append(cards, card)
	}

	if err != nil {
		fmt.Println("Fetching cards fail: ", err)
	} else {
		fmt.Println("Success!")
	}
	latitude, err3 := strconv.ParseFloat(vars["latitude"], 64)
	longitude, err4 := strconv.ParseFloat(vars["longitude"], 64)

	if err3 != nil || err4 != nil {
		fmt.Println("Invalid Params")
		return
	}

	fmt.Println(latitude, longitude)

	helper.GetTopCards(&result, cards, latitude, longitude)

	json.NewEncoder(w).Encode(result)
}
