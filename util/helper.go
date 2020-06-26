package util

import (
	"encoding/json"
	"fmt"
	"strconv"

	model "../models"
	"github.com/bradfitz/slice"
	geo "github.com/kellydunn/golang-geo"
	"github.com/olivere/elastic"
)

// return a geoPoint given latitude and longitude
func getGeoPoint(latitude float64, longitude float64) *geo.Point {
	return geo.NewPoint(latitude, longitude)
}

//Validate the query
func ValidateQuery(searchSource *elastic.SearchSource) bool {

	fmt.Println("Validating Query....")

	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("Error occured during marshalling of query", err1, err2)
		return false
	}
	fmt.Println("Query Validated = ", string(queryJs))
	return true

}

// func Validate User Input
func ValidateInput(strlatitude string, strlongitude string) (float64, float64, bool) {

	fmt.Println("Validating Params...")
	fmt.Println("Latitude: "+strlatitude, "Longitude: "+strlongitude)

	latitude, err1 := strconv.ParseFloat(strlatitude, 64)
	longitude, err2 := strconv.ParseFloat(strlongitude, 64)

	if err1 != nil || err2 != nil {
		return latitude, longitude, false
	}

	fmt.Println("Latitude:", latitude, "Longitude:", longitude)
	fmt.Println("Valid Params!")

	return latitude, longitude, true
}

// returns top 10 Users at least distance from the given Users
func GetTopTenUsers(result *[]model.User, users []model.User, longitude float64, latitude float64) {

	p := getGeoPoint(latitude, longitude)
	slice.Sort(users[:], func(i, j int) bool {
		point1 := getGeoPoint(users[i].Latitude, users[i].Longitude)
		point2 := getGeoPoint(users[j].Latitude, users[j].Longitude)
		dist1 := p.GreatCircleDistance(point1)
		dist2 := p.GreatCircleDistance(point2)
		return dist1 < dist2
	})
	for i := 0; i < 10; i++ {
		*result = append(*result, users[i])
	}
	return
}

// returns top 10 Cards at least distance from the given Users
func GetTopTenCards(result *[]model.Card, cards []model.Card, longitude float64, latitude float64) {
	p := getGeoPoint(latitude, longitude)
	slice.Sort(cards[:], func(i, j int) bool {
		point1 := getGeoPoint(cards[i].Latitude, cards[i].Longitude)
		point2 := getGeoPoint(cards[j].Latitude, cards[j].Longitude)
		dist1 := p.GreatCircleDistance(point1)
		dist2 := p.GreatCircleDistance(point2)
		return dist1 < dist2
	})
	for i := 0; i < 10; i++ {
		*result = append(*result, cards[i])
	}
	return
}
