package busRequest

import (
    "net/http"
    
    "github.com/gorilla/mux"
    "github.com/BryannYeap/take_home_assignment/busTimingService"
)

func BusLineWithBusStopsRequest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    busLineWithBusStops := getBusLineWithBusStopsWithoutForecast(params["id"])
    getEncoder(w).Encode(busLineWithBusStops)
}

func getBusLineWithBusStopsWithoutForecast(busLineID string) busTimingService.BusLineWithBusStops {
    return getBusLineWithBusStops(busLineID, false)
}
