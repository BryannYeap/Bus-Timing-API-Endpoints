package busRequest

import (
    "fmt"
    "os"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/BryannYeap/take_home_assignment/busTimingService"
)

func BusTimingRequest(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    busTiming := getBusTiming(params["id"])
    getEncoder(w).Encode(busTiming)
}

func getBusTiming(busID string) busTimingService.BusTiming {
    var bus busTimingService.Bus

    busFound := false
    currentBuses := getCurrentBuses()
    busIDInt := convertStringToInt(busID)

    for indexOfBus := range currentBuses.Buses {
        bus = currentBuses.Buses[indexOfBus]
        if bus.Vehicle_ID == busIDInt {
            busFound = true
            break;
        }
    }

    if !busFound {
        fmt.Println("Bus not found")
        os.Exit(1)
    }

    return instantiateBusTimingUsingBus(bus)
}

func instantiateBusTimingUsingBus(bus busTimingService.Bus) busTimingService.BusTiming {
    //var busTimingService.busTiming

    //busId := bus.Vehicle_ID
    //busLines := bus.BusLines
    var route []busTimingService.Route 

    return busTimingService.BusTiming{3, route}
}
