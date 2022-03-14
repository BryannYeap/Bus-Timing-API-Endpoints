package busTimingService

// Response to BusStopRequest
type BusStop struct {
	BusStop_ID int
	Name string
	BusForecasts []BusForecast `json:",omitempty"`
}

// Response to BusLineRequest
type BusLineWithBuses struct {
	BusLine BusLine
	Buses []Bus
}

// Response to BusLineWithBusStopsRequest
type BusLineWithBusStops struct {
	BusLine BusLine
	BusStops []BusStop
}

// Response to CurrentBusesRequest
type CurrentBuses struct {
	Buses []BusWithBusLines 
}

// Response to BusTimingRequest
type BusTiming struct {
	Bus Bus
	BusLines []BusLineWithBusStops
}
