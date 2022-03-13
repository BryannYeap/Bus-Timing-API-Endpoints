package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/BryannYeap/take_home_assignment/busRequest"
)

func main() {
    handleRequests()
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/", busRequest.HomePageRequest)
    myRouter.HandleFunc("/busstop/{id}", busRequest.BusStopRequest)
    myRouter.HandleFunc("/busline/{id}", busRequest.BusLineRequest)
    myRouter.HandleFunc("/busline_with_busstops/{id}", busRequest.BusLineWithBusStopsRequest)
    myRouter.HandleFunc("/currentbuses", busRequest.CurrentBusesRequest)
    myRouter.HandleFunc("/bustiming/{id}", busRequest.BusTimingRequest)

    log.Fatal(http.ListenAndServe(":4000", myRouter))
}
