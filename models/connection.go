package models

import (
	elastic "github.com/olivere/elastic"
)

func MakeDBConnection() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	return client, err
}
