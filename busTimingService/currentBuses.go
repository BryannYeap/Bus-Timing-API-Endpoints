package busTimingService

type CurrentBuses struct {
	Buses []Bus 
}

type Bus struct {
	Vehicle_ID int
	BusLines []BusLine
}

type BusLine struct {
	RV_ID int
	Name string
}
