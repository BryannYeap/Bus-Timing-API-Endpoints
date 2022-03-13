package busRequest

import (
    "net/http"

    "github.com/gorilla/mux"
    "github.com/BryannYeap/take_home_assignment/busTimingService"
)

func BusTimingRequest(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    encoder := getEncoder(w)
    busTiming, err := getBusTiming(params["bus_vehicleid"])

    if err != nil {
        encoder.Encode(BusRequestError{ErrorMessage: err.Error()})
    } else {
        encoder.Encode(busTiming)
    }
}

func getBusTiming(busID string) (busTimingService.BusTiming, error) {
    var busWithBusLines busTimingService.BusWithBusLines

    busFound := false
    currentBuses := getCurrentBuses()
    busIDInt, convertErr := convertStringToInt(busID)

    for indexOfBus := range currentBuses.Buses {
        busWithBusLines = currentBuses.Buses[indexOfBus]
        bus := busWithBusLines.Bus

        if bus.Vehicle_ID == busIDInt {
            busFound = true
            break;
        }
    }

    if convertErr != nil || !busFound {
        return busTimingService.BusTiming{}, invalidVehicleIDError
    }

    return instantiateBusTimingUsingBus(busWithBusLines), nil
}

func instantiateBusTimingUsingBus(busWithBusLines busTimingService.BusWithBusLines) busTimingService.BusTiming {
    bus := busWithBusLines.Bus
    busLinesWithBusStops := []busTimingService.BusLineWithBusStops{}

    busLines := busWithBusLines.BusLines
    for _, busLine := range busLines {
        busLineWithBusStops, _ := getBusLineWithBusStops(convertIntToString(busLine.RV_ID), true)
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
                        BusLine: busForecast.BusLine,
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
