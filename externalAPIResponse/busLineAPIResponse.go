package externalAPIResponse

type BusLineAPIResponse struct {
	ID int
	Name string
	Vehicles []Vehicle
}

type Vehicle struct {
	Vehicle_ID int
}
