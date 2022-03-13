package busRequest

import (
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    "github.com/BryannYeap/take_home_assignment/busTimingService"
    "github.com/BryannYeap/take_home_assignment/externalAPIResponse"
)

func BusLineRequest(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    encoder := getEncoder(w)
    _, convertErr := convertStringToInt(params["busline_id"])
    busLine, err := getBusLine(params["busline_id"])

    if convertErr != nil || err != nil {
        encoder.Encode(BusRequestError{ErrorMessage: invalidBusLineIDError.Error()})
    } else {
        encoder.Encode(convertBusLineAPIResponseObjectToBusLine(busLine))
    }
}

func getBusLine(busLineID string) (externalAPIResponse.BusLineAPIResponse, error) {
    const getBusLineURL = "https://dummy.uwave.sg/busline/"
    var busLineAPIResponseObject externalAPIResponse.BusLineAPIResponse

    content := performGetRequest(getBusLineURL + busLineID)

    jsonErr := json.Unmarshal(content, &busLineAPIResponseObject)
    if jsonErr != nil {
        return externalAPIResponse.BusLineAPIResponse{}, invalidBusLineIDError
    }

    return busLineAPIResponseObject, nil
}

func convertBusLineAPIResponseObjectToBusLine(
        busLineAPIResponseObject externalAPIResponse.BusLineAPIResponse) busTimingService.BusLineWithBuses {

    buses := []busTimingService.Bus{}

    for _, vehicle := range busLineAPIResponseObject.Vehicles {
        buses = append(buses, busTimingService.Bus{
            Vehicle_ID: vehicle.Vehicle_ID,
        })
    }

    return busTimingService.BusLineWithBuses{
        BusLine: busTimingService.BusLine{
            RV_ID: busLineAPIResponseObject.ID,
            Name: busLineAPIResponseObject.Name,
        },
        Buses: buses,
    }
}
