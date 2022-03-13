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
    var busWithBusLines busTimingService.BusWithBusLines

    busFound := false
    currentBuses := getCurrentBuses()
    busIDInt := convertStringToInt(busID) // <- no need own line

    for indexOfBus := range currentBuses.Buses {
        busWithBusLines = currentBuses.Buses[indexOfBus]
        bus := busWithBusLines.Bus

        if bus.Vehicle_ID == busIDInt {
            busFound = true
            break;
        }
    }

    if !busFound {
        fmt.Println("Bus not found")
        os.Exit(1)
    }

    return instantiateBusTimingUsingBus(busWithBusLines)
}

func instantiateBusTimingUsingBus(busWithBusLines busTimingService.BusWithBusLines) busTimingService.BusTiming {
    //var busTiming busTimingService.BusTiming
    /**
    bus := busWithBusLines.Bus
    busLines := busWithBusLines.BusLines
    busId := bus.Vehicle_ID
    **/
    var route []busTimingService.Route 

    return busTimingService.BusTiming{3, route}
}
