package busTimingService

type BusStop struct {
	BusStop_ID int
	Name string
	BusForecasts []BusForecast `json:",omitempty"`
}

type BusLine struct {
	RV_ID int
	Name string
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
