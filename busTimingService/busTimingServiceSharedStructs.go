package busTimingService

type BusLine struct {
	RV_ID int
	Name string
}

type BusStop struct {
	BusStop_ID int
	Name string
	BusForecasts []BusForecast `json:",omitempty"`
}

type BusForecast struct {
	Bus Bus
	Forecast_In_Seconds float64
	Forecast_In_Minutes float64
	BusLine BusLine `json:",omitempty"`
}

type Bus struct {
	Vehicle_ID int
}

type BusWithBusLines struct {
	Bus Bus
	BusLines []BusLine
}
