package main

import (
  "fmt"
  "strconv"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "encoding/json"
  "github.com/gorilla/mux"
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

func main() {
    handleRequests()
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/busstop/{id}", busStop)
    myRouter.HandleFunc("/busline/{id}", busLine)
    myRouter.HandleFunc("/currentbuses", currentBuses)

    log.Fatal(http.ListenAndServe(":4000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html> Navigate here to see API end point documentation: " +
        "<a href='https://github.com/BryannYeap/take_home_assignment'> Bus Timing API Endpoints </a></html>")
}

func busStop(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    busStop := getBusStop(params["id"])
    getEncoder(w).Encode(busStop)
}

func getBusStop(id string) externalAPIResponse.BusStopAPIResponse {
    const getBusStopURL = "https://dummy.uwave.sg/busstop/"
    var busStopResponseObject externalAPIResponse.BusStopAPIResponse
    
    content := performGetRequest(getBusStopURL + id)

    jsonErr := json.Unmarshal(content, &busStopResponseObject)
    if jsonErr != nil {
        fmt.Println(jsonErr)
        os.Exit(1)
    }

    return busStopResponseObject
    // fmt.Println(responseString.String())
    //fmt.Printf("%+v\n", busStopResponseObject);  
}

func performGetRequest(url string) []byte {
    response, err := http.Get(url)
    if err != nil {
    fmt.Println(err)
        os.Exit(1)
    }

    defer response.Body.Close()

    //var responseString strings.Builder
    content, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    //responseString.Write(content)

    return content
}

func getEncoder(w http.ResponseWriter) *json.Encoder  {
    encoder := json.NewEncoder(w)
    encoder.SetEscapeHTML(false)
    encoder.SetIndent("", "    ")
    return encoder
}

func busLine(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    busLine := getBusLine(params["id"])
    getEncoder(w).Encode(busLine)
}

func getBusLine(id string) externalAPIResponse.BusLineAPIResponse {
    const getBusLineURL = "https://dummy.uwave.sg/busline/"
    var busLineResponseObject externalAPIResponse.BusLineAPIResponse

    content := performGetRequest(getBusLineURL + id)

    jsonErr := json.Unmarshal(content, &busLineResponseObject)
    if jsonErr != nil {
        fmt.Println(jsonErr)
        os.Exit(1)
    }

    return busLineResponseObject
    // fmt.Println(responseString.String())
    //fmt.Printf("%+v\n", busStopResponseObject);  
}

func currentBuses(w http.ResponseWriter, r *http.Request) {
    getEncoder(w).Encode(getCurrentBuses())
}

func getCurrentBuses() busTimingService.CurrentBuses {
    var busLineAPIResponse externalAPIResponse.BusLineAPIResponse
    var buses []busTimingService.Bus

    busIDToBusLineIDsMap := make(map[int][]string)
    busLineIDtoNameMap := make(map[string]string)

    for _, busLineID := range busLineIDs {
        busLineAPIResponse = getBusLine(busLineID)
        _, exists := busLineIDtoNameMap[busLineID]
        if !exists {
            busLineIDtoNameMap[busLineID] = busLineAPIResponse.Name
        }

        for _, bus := range busLineAPIResponse.Vehicles {
            busIDToBusLineIDsMap[bus.Vehicle_ID] = append(busIDToBusLineIDsMap[bus.Vehicle_ID], busLineID)
        }
    }

    for busID, busLineIDs := range busIDToBusLineIDsMap {
        var busLines []busTimingService.BusLine

        for _, busLineID := range busLineIDs {
            busLineIDInt, _ := strconv.Atoi(busLineID)
            busLines = append(busLines, busTimingService.BusLine{
                RV_ID: busLineIDInt,
                Name: busLineIDtoNameMap[busLineID],
            })
        }

        buses = append(buses, busTimingService.Bus{
            Vehicle_ID: busID,
            BusLines: busLines,
        })
    }

    return busTimingService.CurrentBuses{Buses: buses}
}
