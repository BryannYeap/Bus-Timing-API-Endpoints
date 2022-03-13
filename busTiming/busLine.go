package busTiming

type BusLine struct {
  ID int `json"id"`
  Name string `json"string"`
  RouteName string `json"string"`
  Vehicles: Buses `json"vehicles"`
  BusStops BusStops `json"-"`
}
