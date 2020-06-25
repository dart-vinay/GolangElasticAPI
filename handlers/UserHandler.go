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

func GetUsersForSubDistrictCode(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Inside GetUsersForSubDistrictCode method ")

	vars := mux.Vars(r)
	subDistrictPrefix := vars["subDistrictCode"]
	var users []model.User

	numberOfRecords := 50

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMultiMatchQuery(subDistrictPrefix, "subDistrictCode").Type("phrase_prefix"))

	// To verify the query
	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}
	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))

	searchService := ESClientGlobal.Search().Index("users").SearchSource(searchSource).Size(numberOfRecords)
	searchResult, err := searchService.Do(ctxGlobal)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var user model.User
		err := json.Unmarshal(hit.Source, &user)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}
		users = append(users, user)
	}

	if err != nil {
		fmt.Println("Fetching student fail: ", err)
	} else {
		for _, s := range users {
			fmt.Printf("User found userID: %s, Latitude: %f, Longitude: %f \n", s.User, s.Longitude, s.Latitude)
		}
	}

	json.NewEncoder(w).Encode(users)

}

func GetTopUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside GetTopUsers method")

	vars := mux.Vars(r)
	var users []model.User
	var result []model.User
	numberOfRecords := 200

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchAllQuery())

	// To verify the query
	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}
	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))

	searchService := ESClientGlobal.Search().Index("users").SearchSource(searchSource).Size(numberOfRecords)
	searchResult, err := searchService.Do(ctxGlobal)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var user model.User
		err := json.Unmarshal(hit.Source, &user)
		if err != nil {
			fmt.Println("[Getting Users][Unmarshal] Err=", err)
		}
		users = append(users, user)
	}

	if err != nil {
		fmt.Println("Fetching users fail: ", err)
	} else {
		fmt.Println("Success!")
	}
	latitude, err3 := strconv.ParseFloat(vars["latitude"], 64)
	longitude, err4 := strconv.ParseFloat(vars["longitude"], 64)

	if err3 != nil || err4 != nil {
		fmt.Println("Invalid Params")
		return
	}

	helper.GetTopUsers(&result, users, latitude, longitude)

	json.NewEncoder(w).Encode(result)
}

func UpdateUserLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside UpdateUserLocation method")
}
