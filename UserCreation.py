# Run the command `python3 UserCreation.py` from the project root directory for populating users data


import json
import os
import pandas as pd
import sys

from datetime import datetime, date

host = "localhost"
port = "9200"
apiEndPoint = "/users/_doc"


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

    userName = "User"
    subDistrictPrefix = "UP_LA_"
    for ind in range(latLongData.shape[0]):

        district = subDistrictData.iloc[ind,0].strip()
        district = district.split(" ")
        district = "_".join(district)

        print("T".join(str(datetime.now()).split(" ")))

        id = userName + "_"  + str(ind+1)
        latitude = latLongData.iloc[ind, 0]
        longitude = latLongData.iloc[ind, 1]
        subDistrictCode = subDistrictPrefix+district

        json_str = {
            "user": id,
            "createdAt": "2018-09-22T12:42:31Z",
            "latitude": latitude,
            "longitude": longitude,
            "subDistrictCode": subDistrictCode
        }
        # print(json_str)
        command = "curl -X POST -H 'Content-type: application/json' --data '" + json.dumps(json_str, default=myconverter) + "' " + apiToHit
        os.system(command)


if __name__ == "__main__":
    generate()
