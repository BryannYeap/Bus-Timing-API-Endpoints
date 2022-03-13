package busRequest

import (
    "net/http"

    "github.com/BryannYeap/take_home_assignment/externalAPIResponse"
    "github.com/BryannYeap/take_home_assignment/busTimingService"
)

func CurrentBusesRequest(w http.ResponseWriter, r *http.Request) {
    getEncoder(w).Encode(getCurrentBuses())
}

func getCurrentBuses() busTimingService.CurrentBuses {
    var busLineAPIResponse externalAPIResponse.BusLineAPIResponse
    var buses []busTimingService.Bus

    busIDToBusLinesMap := make(map[int][]busTimingService.BusLine)

    for _, busLineID := range busLineIDs {
        busLineAPIResponse = getBusLine(busLineID)

        for _, bus := range busLineAPIResponse.Vehicles {
            busIDInt := bus.Vehicle_ID
            rvID := busLineAPIResponse.ID
            routeName := busLineAPIResponse.Name

            busIDToBusLinesMap[busIDInt] = append(busIDToBusLinesMap[busIDInt], busTimingService.BusLine{
                RV_ID: rvID,
                Name: routeName,
            })
        }
    }

    for busID, busLines := range busIDToBusLinesMap {
        buses = append(buses, busTimingService.Bus{
            Vehicle_ID: busID,
            BusLines: busLines,
        })
    }

    return busTimingService.CurrentBuses{Buses: buses}
}
