# Users Index
{
	"settings": {
    	"number_of_shards": 1,
    	"number_of_replicas": 1
	},
   "mappings": {
       "properties": {
         "user": {
               "type": "text"
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
               "type": "text"
         }
     }
   }
}

#Cards Index
{
	"settings": {
    	"number_of_shards": 1,
    	"number_of_replicas": 1
	},
   "mappings": {
       "properties": {
         "cardID": {
               "type": "text"
         },
         "createdAt": {
               "type": "date"
         },
         "Title": {
               "type": "text"
         },
         "latitude": {
               "type": "float"      
         },
         "longitude": {
               "type": "float"
         },
         "subDistrictCode": {
               "type": "text"
         },
         "createdBy": {
               "type": "text"
         }
     }
   }
}