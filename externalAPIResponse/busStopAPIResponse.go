package externalAPIResponse

type BusStopAPIResponse struct {
  ExternalID int `json"external_id"`
  Forecasts []Forecast `json"forecast"`
}

type Forecast struct {
  ForecastInSeconds float64 `json"forecast_seconds"`
  Route Route `json"route"`
  RouteVariantID int `json"rv_id"`
  VehicleID int `json"vehicle_id"`
}

type Route struct {
  ID int `json"id"`
  Name string `json"name"`
}

