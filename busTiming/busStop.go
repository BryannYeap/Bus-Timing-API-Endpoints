package busTiming

type BusStop struct {
  ExternalID int `json"external_id"`
  ID int `json"id"`
  Name string `json"name"`
  BusForecasts BusForecasts `json"forecast"`
}

type BusStops BusStop[]
