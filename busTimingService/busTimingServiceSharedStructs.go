package busTimingService

type BusLine struct {
	RV_ID int
	Name string
}

type BusForecast struct {
	Bus Bus
	Forecast_In_Seconds float64
	Forecast_In_Minutes float64
	BusLine BusLine
}

type Bus struct {
	Vehicle_ID int
}

type BusWithBusLines struct {
	Bus Bus
	BusLines []BusLine
}
