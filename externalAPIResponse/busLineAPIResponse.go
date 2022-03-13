package externalAPIResponse

type BusLineAPIResponse struct {
	ExternalID string `json"external_id"`
	ID int `json"id"`
	Name string `json"name"`
	Vehicles []Vehicle `json"vehicles"`
}

type Vehicle struct {
	Vehicle_ID int `json"vehicle_id"`
}
