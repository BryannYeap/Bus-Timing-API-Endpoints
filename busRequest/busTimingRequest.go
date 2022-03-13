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

    for indexOfBus := range currentBuses.Buses {
        busWithBusLines = currentBuses.Buses[indexOfBus]
        bus := busWithBusLines.Bus

        if bus.Vehicle_ID == convertStringToInt(busID) {
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
    bus := busWithBusLines.Bus
    busLinesWithBusStops := []busTimingService.BusLineWithBusStops{}
    
    busLines := busWithBusLines.BusLines
    for _, busLine := range busLines {
        busLineWithBusStops := getBusLineWithBusStops(convertIntToString(busLine.RV_ID), true)
        for _, busStop := range busLineWithBusStops.BusStops {
            newBusForecasts := []busTimingService.BusForecast{}
            busForecasts := busStop.BusForecasts
                for _, busForecast := range busForecasts {
                busInBusForecast := busForecast.Bus
                if busInBusForecast == bus {
                    newBusForecasts = append(newBusForecasts, busTimingService.BusForecast{
                        Bus: busForecast.Bus,
	                    Forecast_In_Seconds: busForecast.Forecast_In_Seconds,
	                    Forecast_In_Minutes: busForecast.Forecast_In_Minutes,
                    })
                } 
            }
            busStop.BusForecasts = newBusForecasts
        }
        busLinesWithBusStops = append(busLinesWithBusStops, busLineWithBusStops)
    }

    return busTimingService.BusTiming{
        Bus: bus,
        BusLines: busLinesWithBusStops,
    }
}
