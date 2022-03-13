package busRequest

import (
    "net/http"
    
    "github.com/gorilla/mux"
    "github.com/BryannYeap/take_home_assignment/busTimingService"
)

func BusLineWithBusStopsRequest(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    encoder := getEncoder(w)
    busLineWithBusStops, err := getBusLineWithBusStopsWithoutForecast(params["busline_id"])

    if err != nil {
        encoder.Encode(BusRequestError{ErrorMessage: err.Error()})
    } else {
        encoder.Encode(busLineWithBusStops)
    }
}

func getBusLineWithBusStopsWithoutForecast(busLineID string) (busTimingService.BusLineWithBusStops, error) {
    return getBusLineWithBusStops(busLineID, false)
}
