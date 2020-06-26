# Users Index
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

#Cards Index
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