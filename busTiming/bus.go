package busTiming

type Bus struct {
  VehicleID int `json"vehicle_id"`
  RouteVariantID int `json"routevariant_id"`
}

type Buses []Bus
