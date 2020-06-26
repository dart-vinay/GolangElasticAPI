package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olivere/elastic"

	model "../models"
)

var ESClientGlobal *elastic.Client

var ctxGlobal context.Context

func createDBConnection() *elastic.Client {

	ctx := context.Background()
	ctxGlobal = ctx

	fmt.Println("Establishing ElasticSearch Connection...")

	esClient, err := model.MakeDBConnection()

	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	fmt.Println("ElasticSearch Connection Established!")

	return esClient
}

func HandleRequests() {

	ESClientGlobal = createDBConnection()

	fmt.Println("--------Server Started-----------")

	myRouter := mux.NewRouter().StrictSlash(true)

	//RequestMappings
	myRouter.HandleFunc("/allCardsForUser/{userID}/", GetAllCardsForUser).Methods("GET")
	myRouter.HandleFunc("/allUsersForSDCode/{subDistrictCode}", GetUsersForSubDistrictCode).Methods("GET")
	myRouter.HandleFunc("/topUsers/{latitude}/{longitude}", GetTopUsers).Methods("GET")
	myRouter.HandleFunc("/topCards/{latitude}/{longitude}", GetTopCards).Methods("GET")
	myRouter.HandleFunc("/changePosition/{userID}/{latitude}/{longitude}", UpdateUserLocation).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", myRouter))

}
