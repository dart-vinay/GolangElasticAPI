# Run the command `python3 CardsCreation.py` from the project root directory for populating cards data

import json
import os
import pandas as pd
import sys
import random

from datetime import datetime, date

host = "localhost"
port = "9200"
apiEndPoint = "/cards/doc"


def myconverter(o):
    if isinstance(o, datetime):
        return o.__str__()

def generate():
    pathForData = "Data/LatLongData.xlsx"
    pathForSubDistrictData  = "Data/SubDistrict.xlsx"
    latLongData = pd.read_excel(pathForData)
    subDistrictData = pd.read_excel(pathForSubDistrictData)

    # API to hit the post request
    apiToHit = host + ":" + port + apiEndPoint

    cardID = "Card"
    subDistrictPrefix = "UP_LA_"
    userName = "User_"
    defaultTitle = "Title_"
    for ind in range(latLongData.shape[0]):

        district = subDistrictData.iloc[ind,0].strip()
        district = district.split(" ")
        district = "_".join(district)

        print("T".join(str(datetime.now()).split(" ")))

        id = cardID + "_"  + str(ind+1)
        # createdAt = "T".join(str(datetime.now()).split(" "))
        title = defaultTitle+str(ind+1)
        latitude = latLongData.iloc[ind, 0]
        longitude = latLongData.iloc[ind, 1]
        subDistrictCode = subDistrictPrefix+district
        createdBy = userName+str(random.randint(1,199))
        
        json_str = {
            "cardId": id,
            "createdAt": "2020-06-25T12:42:31Z",
            "title" : title,
            "latitude": latitude,
            "state" : "PUBLISH",
            "longitude": longitude,
            "subDistrictCode": subDistrictCode,
            "createdBy": createdBy
        }
        # print(json_str)
        command = "curl -X POST -H 'Content-type: application/json' --data '" + json.dumps(json_str, default=myconverter) + "' " + apiToHit
        os.system(command)


if __name__ == "__main__":
    generate()
