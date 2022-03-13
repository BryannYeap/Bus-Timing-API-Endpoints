package busRequest

import (
    "fmt"
    "os"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    "github.com/BryannYeap/take_home_assignment/externalAPIResponse"
)

func BusLineRequest(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    busLine := getBusLine(params["id"])
    getEncoder(w).Encode(busLine)
}

func getBusLine(busLineID string) externalAPIResponse.BusLineAPIResponse {
    const getBusLineURL = "https://dummy.uwave.sg/busline/"
    var busLineResponseObject externalAPIResponse.BusLineAPIResponse

    content := performGetRequest(getBusLineURL + busLineID)

    jsonErr := json.Unmarshal(content, &busLineResponseObject)
    if jsonErr != nil {
        fmt.Println(jsonErr)
        os.Exit(1)
    }

    return busLineResponseObject
}