package busTimingService

type BusTiming struct {
	Bus_ID int
	Routes []Route
}

type Route struct {
	BusLine BusLine
	//BusForecast BusForecast
}
