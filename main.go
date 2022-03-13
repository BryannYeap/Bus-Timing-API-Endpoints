package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/BryannYeap/take_home_assignment/externalAPIResponse"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html> Navigate here to see API end point documentation: " +
        "<a href='https://github.com/BryannYeap/take_home_assignment'> Bus Timing API Endpoints </a></html>")
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

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/busstop/{id}", busStop)
    myRouter.HandleFunc("/busline/{id}", busLine)
    log.Fatal(http.ListenAndServe(":4000", myRouter))
}

func getEncoder(w http.ResponseWriter) *json.Encoder  {
    encoder := json.NewEncoder(w)
    encoder.SetEscapeHTML(false)
    encoder.SetIndent("", "    ")
    return encoder
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

func main() {
    handleRequests()
}
