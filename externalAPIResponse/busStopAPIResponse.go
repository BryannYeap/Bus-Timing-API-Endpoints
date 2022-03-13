package externalAPIResponse

type BusStopAPIResponse struct {
	Forecast []Forecast
	ID int
	Name string
}

type Forecast struct {
	Forecast_Seconds float64
	Route Route
	RV_ID int
	Vehicle_ID int
}

type Route struct {
	ID int
	Name string
}

