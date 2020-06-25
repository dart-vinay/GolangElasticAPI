# GolangElasticAPI

## About

The repository aims to provide interation of Golang with ElasticSearch DB. The thought used in building the APIs is based on providing relevant NewsBlocks (referred as `cards` in this project) for a user based on his/her geo-location.

This repository has four main folders and a few data generations scripts:

* **Data** Contains all the raw data we need to populate our elastic search indexes
* **handlers** Contains the API definition and corresponding request handlers
* **models** Contains the structural objects used in this project and DB connection details
* **util** Contains all the helper function that we need for functioning of our services

## Installation

Install Golang and ElasticSearch

## Usage

* Clone the repository
* Run the elastic server(`Port: 9200`) and run the data creation scripts UserCreation.py and CardsCreation.py for data generation
* `Go` ahead and run the following command from the project root directory:

```bash
go run main.go
```

To test the behavior of out we have four major APIs that we have created:
- Get all the users belonging to the particular sub-district (prefix query)
- Get all the cards for a particular user
- Get top relevant cards for the particular user based on geo-location
- Get nearest users for the particular user

Go ahead and test it and let me know your thoughts over this.

## Updates

Will be grooming this repository more with further developments. Cheers!
