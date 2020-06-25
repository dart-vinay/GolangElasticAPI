package util

import (
	model "../models"
	"github.com/bradfitz/slice"
	geo "github.com/kellydunn/golang-geo"
)

// returns top 10 Users at least distance from the given Users
func GetTopUsers(result *[]model.User, users []model.User, longitude float64, latitude float64) {
	p := geo.NewPoint(latitude, longitude)
	slice.Sort(users[:], func(i, j int) bool {
		point1 := geo.NewPoint(users[i].Latitude, users[i].Longitude)
		point2 := geo.NewPoint(users[j].Latitude, users[j].Longitude)
		// find the great circle distance between them
		dist1 := p.GreatCircleDistance(point1)
		dist2 := p.GreatCircleDistance(point2)
		return dist1 > dist2
	})
	for i := 0; i < 10; i++ {
		*result = append(*result, users[i])
	}
	return
}

// returns top 10 Cards at least distance from the given Cards
func GetTopCards(result *[]model.Card, cards []model.Card, longitude float64, latitude float64) {
	p := geo.NewPoint(latitude, longitude)
	slice.Sort(cards[:], func(i, j int) bool {
		point1 := geo.NewPoint(cards[i].Latitude, cards[i].Longitude)
		point2 := geo.NewPoint(cards[j].Latitude, cards[j].Longitude)
		// find the great circle distance between them
		dist1 := p.GreatCircleDistance(point1)
		dist2 := p.GreatCircleDistance(point2)
		return dist1 > dist2
	})
	for i := 0; i < 10; i++ {
		*result = append(*result, cards[i])
	}
	return
}
