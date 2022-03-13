package busRequest

import (
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    "github.com/BryannYeap/take_home_assignment/busTimingService"
    "github.com/BryannYeap/take_home_assignment/externalAPIResponse"
)

func BusStopRequest(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    encoder := getEncoder(w)
    _, convertErr := convertStringToInt(params["id"])
    busStop, err := getBusStop(params["id"])

    if convertErr != nil || err != nil {
        encoder.Encode(BusRequestError{ErrorMessage: invalidBusStopIDError.Error()})
    } else {
        encoder.Encode(convertBusStopAPIResponseObjectToBusStop(busStop))
    }
}

func getBusStop(busStopID string) (externalAPIResponse.BusStopAPIResponse, error) {
    const getBusStopURL = "https://dummy.uwave.sg/busstop/"
    var busStopAPIResponseObject externalAPIResponse.BusStopAPIResponse
    
    content := performGetRequest(getBusStopURL + busStopID)

    jsonErr := json.Unmarshal(content, &busStopAPIResponseObject)
    if jsonErr != nil {
        return externalAPIResponse.BusStopAPIResponse{}, invalidBusStopIDError
    }

    return busStopAPIResponseObject, nil
}

func convertBusStopAPIResponseObjectToBusStop(
        busStopAPIResponseObject externalAPIResponse.BusStopAPIResponse) busTimingService.BusStop {

    busForecasts := []busTimingService.BusForecast{}

    for _, forecast := range busStopAPIResponseObject.Forecast {
        busLineAPIResponse, _ := getBusLine(convertIntToString(forecast.RV_ID))
        busForecasts = append(busForecasts, busTimingService.BusForecast{
            Bus: busTimingService.Bus{Vehicle_ID: forecast.Vehicle_ID},
            Forecast_In_Seconds: forecast.Forecast_Seconds,
	        Forecast_In_Minutes: forecast.Forecast_Seconds / 60,
	        BusLine: busTimingService.BusLine{
                RV_ID: busLineAPIResponse.ID,
                Name: busLineAPIResponse.Name,
            },
        })
    }

    return busTimingService.BusStop{
        BusStop_ID: busStopAPIResponseObject.ID,
        Name: busStopAPIResponseObject.Name,
        BusForecasts: busForecasts,
    }
}
