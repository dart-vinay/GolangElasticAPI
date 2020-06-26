# GolangElasticAPI

## About

The repository aims to provide interation of Golang with ElasticSearch DB. The thought used in building the APIs is based on providing relevant NewsBlocks (referred as `cards` in this project) for a user based on his/her geo-location.

This repository has four main folders and a few data generation scripts:

* **Data** Contains all the raw data we need to populate our elastic search indexes
* **handlers** Contains the API definition and corresponding request handlers
* **models** Contains the structural objects used in this project and DB connection details
* **util** Contains all the helper function that we need for functioning of our services

## Installation

Install Golang and ElasticSearch

## Usage

* Clone the repository
* Run the elastic server(`Port: 9200`)
* Create index `users` using the following structure:

```javascript
{
      "settings": {
            "number_of_shards": 1,
            "number_of_replicas": 1,
            "analysis": {
                  "analyzer": {
                        "not_analyzed": {
                              "type": "custom",
                              "tokenizer": "standard"
                        }
                  }
            }
      },
      "mappings": {
            "properties": {
                  "user": {
                        "type": "text",
                        "analyzer": "not_analyzed"
                  },
                  "createdAt": {
                        "type": "date"
                  },
                  "state": {
                        "type": "text"
                  },
                  "latitude": {
                        "type": "float"      
                  },
                  "longitude": {
                        "type": "float"
                  },
                  "subDistrictCode": {
                        "type": "text",
                        "analyzer": "not_analyzed"
                  }
            }
      }
}
```

* Create index `cards` using the following structure:

```javascript
{
      "settings": {
            "number_of_shards": 1,
            "number_of_replicas": 1,
            "analysis": {
                  "analyzer": {
                        "not_analyzed": {
                              "type": "custom",
                              "tokenizer": "standard"
                        }
                  }
            }
      },
      "mappings": {
            "properties": {
                  "cardId": {
                        "type": "text",
                        "analyzer": "not_analyzed"
                  },
                  "createdAt": {
                        "type": "date"
                  },
                  "title": {
                        "type": "text"
                  },
                  "state": {
                        "type": "text"
                  },
                  "latitude": {
                        "type": "float"      
                  },
                  "longitude": {
                        "type": "float"
                  },
                  "subDistrictCode": {
                        "type": "text",
                        "analyzer": "not_analyzed"
                  },
                  "createdBy": {
                        "type": "text",
                        "analyzer": "not_analyzed"
                  }
            }
      }
}
```
* Run the data creation scripts UserCreation.py and CardsCreation.py for data generation
* `Go` ahead and run the following command from the project root directory:

```bash
go run main.go

```
This will start the backend server at `localhost:8080`

To test the behavior of out we have four major APIs that we have created:
- Update user's geo-position
  - endPoint: `/changePosition/{userID}/{latitude}/{longitude}`
- Get all the users belonging to the particular sub-district (prefix query) 
  - endPoint: `/allUsersForSDCode/{subDistrictCode}`
- Get all the cards for a particular user
  - endPoint: `/allCardsForUser/{userID}/`
- Get top relevant cards for the particular user based on geo-position
  - endPoint: `/topCards/{latitude}/{longitude}`
- Get nearest users for the particular user based on geo-position
  - endPoint: `/topUsers/{latitude}/{longitude}`

Go ahead and test it and let me know your thoughts over this.

## Updates

Will be grooming this repository more with further developments. Cheers!
