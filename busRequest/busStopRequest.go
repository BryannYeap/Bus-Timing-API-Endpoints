package busRequest

import (
    "fmt"
    "os"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    "github.com/BryannYeap/take_home_assignment/externalAPIResponse"
)

func BusStopRequest(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    busStop := getBusStop(params["id"])
    getEncoder(w).Encode(busStop)
}

func getBusStop(busStopID string) externalAPIResponse.BusStopAPIResponse {
    const getBusStopURL = "https://dummy.uwave.sg/busstop/"
    var busStopResponseObject externalAPIResponse.BusStopAPIResponse
    
    content := performGetRequest(getBusStopURL + busStopID)

    jsonErr := json.Unmarshal(content, &busStopResponseObject)
    if jsonErr != nil {
        fmt.Println(jsonErr)
        os.Exit(1)
    }

    return busStopResponseObject
}