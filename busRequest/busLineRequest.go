package busRequest

import (
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    "github.com/BryannYeap/take_home_assignment/busTimingService"
    "github.com/BryannYeap/take_home_assignment/externalAPIResponse"
)

func BusLineRequest(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    encoder := getEncoder(w)
    _, convertErr := convertStringToInt(params["id"])
    busLine, err := getBusLine(params["id"])

    if convertErr != nil || err != nil {
        encoder.Encode(BusRequestError{ErrorMessage: invalidBusLineIDError.Error()})
    } else {
        encoder.Encode(convertBusLineAPIResponseObjectToBusLine(busLine))
    }
}

func getBusLine(busLineID string) (externalAPIResponse.BusLineAPIResponse, error) {
    const getBusLineURL = "https://dummy.uwave.sg/busline/"
    var busLineAPIResponseObject externalAPIResponse.BusLineAPIResponse

    content := performGetRequest(getBusLineURL + busLineID)

    jsonErr := json.Unmarshal(content, &busLineAPIResponseObject)
    if jsonErr != nil {
        return externalAPIResponse.BusLineAPIResponse{}, invalidBusLineIDError
    }

    return busLineAPIResponseObject, nil
}

func convertBusLineAPIResponseObjectToBusLine(
        busLineAPIResponseObject externalAPIResponse.BusLineAPIResponse) busTimingService.BusLine {
    return busTimingService.BusLine{
        RV_ID: busLineAPIResponseObject.ID,
        Name: busLineAPIResponseObject.Name,
    }
}