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

func GetUsersForSubDistrictCode(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Inside GetUsersForSubDistrictCode method ")

	vars := mux.Vars(r)
	subDistrictPrefix := vars["subDistrictCode"]
	var users []model.User

	numberOfRecords := 50

	searchSource := elastic.NewSearchSource().Query(elastic.NewMultiMatchQuery(subDistrictPrefix, "subDistrictCode").Type("phrase_prefix"))

	// Validate Query
	isValid := helper.ValidateQuery(searchSource)

	if !isValid {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	searchService := ESClientGlobal.Search().Index("users").SearchSource(searchSource).Size(numberOfRecords)
	searchResult, err := searchService.Do(ctxGlobal)
	if err != nil {
		fmt.Println("InvalidResult", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var user model.User
		err := json.Unmarshal(hit.Source, &user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error unmarshalling", err)
		}
		users = append(users, user)
	}

	if err != nil {
		fmt.Println("Invalid Result", err)
	} else {
		for _, s := range users {
			fmt.Printf("User found userID: %s, Latitude: %f, Longitude: %f \n", s.User, s.Longitude, s.Latitude)
		}
	}
	fmt.Println("Success!")
	json.NewEncoder(w).Encode(users)

}

func GetTopUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside GetTopUsers method")

	vars := mux.Vars(r)
	var users []model.User
	var result []model.User
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

	searchService := ESClientGlobal.Search().Index("users").SearchSource(searchSource).Size(numberOfRecords)
	searchResult, err := searchService.Do(ctxGlobal)
	if err != nil {
		fmt.Println("Invalid Result", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var user model.User
		err := json.Unmarshal(hit.Source, &user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error unmarshalling", err)
		}
		users = append(users, user)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Fetching users fail: ", err)
	} else {
		fmt.Println("Success!")
	}

	helper.GetTopTenUsers(&result, users, longitude, latitude)

	json.NewEncoder(w).Encode(result)
}

func UpdateUserLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside UpdateUserLocation method")

	vars := mux.Vars(r)
	userID := vars["userID"]
	latitude := vars["latitude"]
	longitude := vars["longitude"]

	//validateParams
	_, _, validParams := helper.ValidateInput(latitude, longitude)

	if !validParams {
		fmt.Println("Invalid Params")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	searchSource := elastic.NewSearchSource().Query(elastic.NewMultiMatchQuery(userID, "user"))

	// Validate Query
	isValid := helper.ValidateQuery(searchSource)

	if !isValid {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	scriptString := "ctx._source.longitude = " + longitude + ";" + "ctx._source.latitude = " + latitude

	searchService := ESClientGlobal.UpdateByQuery().Query(elastic.NewMultiMatchQuery(userID, "user")).Index("users").Script(elastic.NewScript(scriptString).Params(map[string]interface{}{"tag": "doc"}))

	_, err := searchService.Do(ctxGlobal)

	if err != nil {
		fmt.Println("Invalid Result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Success!")

}
