package busRequest

import "errors"

type BusRequestError struct {
	ErrorMessage string
}

var invalidBusStopIDError = errors.New("Bus Stop with given ID was not found or is invalid.")
var invalidBusLineIDError = errors.New("Bus Line with given ID was not found or is invalid.")
var invalidVehicleIDError = errors.New("Bus with given Vehicle ID is currently not in service or does not exist.")