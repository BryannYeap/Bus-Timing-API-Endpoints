package busRequest

import (
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    "github.com/BryannYeap/take_home_assignment/externalAPIResponse"
)

func BusStopRequest(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    encoder := getEncoder(w)
    busStop, err := getBusStop(params["id"])

    if err != nil {
        encoder.Encode(BusRequestError{ErrorMessage: err.Error()})
    } else {
        encoder.Encode(busStop)
    }
}

func getBusStop(busStopID string) (externalAPIResponse.BusStopAPIResponse, error) {
    const getBusStopURL = "https://dummy.uwave.sg/busstop/"
    var busStopResponseObject externalAPIResponse.BusStopAPIResponse
    
    content := performGetRequest(getBusStopURL + busStopID)

    jsonErr := json.Unmarshal(content, &busStopResponseObject)
    if jsonErr != nil {
        return externalAPIResponse.BusStopAPIResponse{}, invalidBusStopIDError
    }

    return busStopResponseObject, nil
}