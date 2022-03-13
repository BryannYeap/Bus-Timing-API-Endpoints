package externalAPIResponse

type BusLineAPIResponse struct {
  ExternalID int `json"external_id"`
  ID int `json"id"`
  Name string `json"name"`
  Vehicles []Vehicle `json"vehicles"`
}

type Vehicle struct {
  VehicleID int `json"vehicle_id"`
}
