package busRequest

import (
"fmt"
"os"
    "errors"
    "log"
    "strconv"
    "io/ioutil"
    "net/http"
    "encoding/json"

    "github.com/BryannYeap/take_home_assignment/externalAPIResponse"
    "github.com/BryannYeap/take_home_assignment/busTimingService"
)

var busLineIDs = []string{"44478", "44479", "44480", "44481"}
var busStopIDs = []string{
    "378204", "383050", "378202", "383049", "382998", "378237", "378233", "378230",
    "378229", "378228", "378227", "382995", "378224", "378226", "383010", "383009",
    "383006", "383004", "378234", "383003", "378222", "383048", "378203", "382999",
    "378225", "383014", "383013", "383011", "377906", "383018", "383015", "378207",
}

func performGetRequest(url string) []byte {
    response, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer response.Body.Close()

    content, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    return content
}

func getEncoder(w http.ResponseWriter) *json.Encoder  {
    encoder := json.NewEncoder(w)
    encoder.SetEscapeHTML(false)
    encoder.SetIndent("", "    ")
    return encoder
}

func getBusLineWithBusStops(busLineID string, includeBusForecast bool) (busTimingService.BusLineWithBusStops, error) {
    busStops := []busTimingService.BusStop{}
    
    busLineAPIResponse, getBusLineErr := getBusLine(busLineID)
    busLineIDInt, convertErr := convertStringToInt(busLineID)
    if getBusLineErr != nil || convertErr != nil {
        return busTimingService.BusLineWithBusStops{}, invalidBusLineIDError
    }

    busLine := busTimingService.BusLine{
        RV_ID: busLineAPIResponse.ID,
        Name: busLineAPIResponse.Name,
    }

    for _, busStopID := range busStopIDs {
        busStopAPIResponse, _ := getBusStop(busStopID)
        forecasts := busStopAPIResponse.Forecast

        for _, forecast := range forecasts {
            if forecast.RV_ID == busLineIDInt {
                busStops = append(busStops, instantiateBusStop(busStopAPIResponse, includeBusForecast))
                break;
            }
        }
    }

    return busTimingService.BusLineWithBusStops{
        BusLine: busLine,
        BusStops: busStops,
    }, nil
}

func instantiateBusStop(busStopAPIResponse externalAPIResponse.BusStopAPIResponse,
                        includeBusForecast bool) busTimingService.BusStop {
    
    busStopID := busStopAPIResponse.ID
    name := busStopAPIResponse.Name

    if includeBusForecast {
        return convertBusStopAPIResponseObjectToBusStop(busStopAPIResponse)
    } else {
        return busTimingService.BusStop{
            BusStop_ID: busStopID,
            Name: name,
        }
    }
}

func convertStringToInt(stringToConvert string) (int, error) {
    intToReturn, err := strconv.Atoi(stringToConvert)
    if err != nil {
        return 0, errors.New("Cannot convert given string to integer")
    }
    return intToReturn, nil
}

func convertIntToString(intToConvert int) string {
    return strconv.Itoa(intToConvert)
}