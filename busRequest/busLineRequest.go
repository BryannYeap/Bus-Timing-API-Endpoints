package busRequest

import (
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    "github.com/BryannYeap/take_home_assignment/externalAPIResponse"
)

func BusLineRequest(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    encoder := getEncoder(w)
    busLine, err := getBusLine(params["id"])

    if err != nil {
        encoder.Encode(BusRequestError{ErrorMessage: err.Error()})
    } else {
        encoder.Encode(busLine)
    }
}

func getBusLine(busLineID string) (externalAPIResponse.BusLineAPIResponse, error) {
    const getBusLineURL = "https://dummy.uwave.sg/busline/"
    var busLineResponseObject externalAPIResponse.BusLineAPIResponse

    content := performGetRequest(getBusLineURL + busLineID)

    jsonErr := json.Unmarshal(content, &busLineResponseObject)
    if jsonErr != nil {
        return externalAPIResponse.BusLineAPIResponse{}, invalidBusLineIDError
    }

    return busLineResponseObject, nil
}