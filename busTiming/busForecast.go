package busTiming

type BusForecast struct {
  Bus Bus `json"bus"`
  ForecastInSeconds float64 `json"forecast_seconds"`
  ForecastInMinutes float64 `json"forecast_minutes"`
  
}

type BusForcasts []BusForecast
