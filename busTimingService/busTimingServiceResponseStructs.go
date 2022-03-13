package busTimingService

type BusStop struct {
	BusStop_ID int
	Name string
	BusForecasts []BusForecast `json:",omitempty"`
}

type BusLineWithBuses struct {
	BusLine BusLine
	Buses []Bus
}

type BusLineWithBusStops struct {
	BusLine BusLine
	BusStops []BusStop
}

type CurrentBuses struct {
	Buses []BusWithBusLines 
}

type BusTiming struct {
	Bus Bus
	BusLines []BusLineWithBusStops
}
