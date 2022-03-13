package externalAPIResponse

type BusStopAPIResponse struct {
	ExternalID string `json"external_id"`
	Forecast []Forecast `json"forecast"`
	ID int `json"id"`
	Name string `json"name"`
}

type Forecast struct {
	Forecast_Seconds float64 `json"forecast_seconds"`
	Route Route `json"route"`
	RV_ID int `json"rv_id"`
	Vehicle_ID int `json"vehicle_id"`
}

type Route struct {
	ID int `json"id"`
	Name string `json"name"`
}

