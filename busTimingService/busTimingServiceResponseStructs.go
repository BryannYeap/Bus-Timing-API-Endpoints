package busTimingService

type BusLineWithBusStops struct {
	BusLine BusLine
	BusStops []BusStop
}

type CurrentBuses struct {
	Buses []BusWithBusLines 
}

type BusTiming struct {
	Bus_ID int
	Routes []Route
}
