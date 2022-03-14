## Bus Timing Service API Endpoint Documentation

Given in each section below are the API Endpoints for the Bus Timing Service.

To test out the service, follow the steps below:
  1. Download the zip and extract the code from [here](https://github.com/BryannYeap/take_home_assignment)
  2. Navigate into the take_home_assignment directory and run the command `go build`
  3. An executable file will be created. Run that file. 
      - Alternatively, obtain the executable file from [here](https://github.com/BryannYeap/take_home_assignment/releases/tag/v1.0)
  4. Upon running the file, open your brower and type `"localhost:8080"` followed by any of the API endpoints given below

Note: The curly brackets in the API endpoints are params

### Homepage

HTTP Request: GET

API Endpoint: `"/"`

Params: None

Response: Provides a link that navigates client to the API Endpoint Documentation

### BusStop

HTTP Request: GET

API Endpoint: `"/busstop/{busstop_id}"`

Params: The id of the bus stop to be queried

Response: Gets the id (given) and name of the bus stop. Additionally, if any buses are forcasted to arrive at this bus stop, the forecasts will be visible as well. Else, no forecasts will be shown.

### BusLine

HTTP Request: GET

APT Endpoint: `"/busline/{busline_id}"`

Params: The id of the bus line to be queried

Response: Gets the id (given), the name of the bus line, and all buses that are currently on this bus line service.

### BusLine with BusStops

HTTP Request: GET

APT Endpoint: `"/busline_with_busstops/{busline_id}"`

Params: The id of the bus line to be queried

Response: Gets the id (given), the name of the bus line, and all bus stops that are on the route of this bus line.

### Current Buses

HTTP Request: GET

APT Endpoint: `"/currentbuses"`

Params: None

Response: Gets all buses that are currently in service, as well as the bus lines that they are on.

### Bus Timing

HTTP Request: GET

APT Endpoint: `"/bustiming/{bus_vehicleid}"`

Params: The vehicle id of the bus to be queried

Response: Gets the id (given) of the bus, and all bus lines that this bus is currently serving. Additionally, for every bus line that this bus is serving, the forecasts of __this bus__ for all of the bus stops on the bus line will be retrieved.

### Code Structure

#### Package: externalAPIResponse

This package contains the structs obtained from performing a get request using the given external API. However, the structs only contain fields that I believe are useful for the bus timing service. For instance, fields that were originally present such as locations (i.e. lon and lat), projections, stats, etc were omitted from the structs in externalAPIResponse. A client does not need to know the location of the bus stops and bus lines in order to provide a bus timing, since the forecast was already provided by the external API.

#### Package: busTimingService

This package contains the structs that are used / provided to the clients. The file `busTimingServiceResponseStructs.go` contain all the structs that are returned as responses to HTTP requests of the bus timing service API. The file `busTimingServiceSharedStructs.go` contain all the intermediate or base structs that are used as fields in the response structs mentioned earlier.

#### Packge: busRequest

This package contain files that each encapsulate a HTTP request of the bus timing service API. Each file contains functions relevant to the request that it encapsulates.

### Program Flow

Here is an example of the logical flow of the program:
    1. An API endpoint will be queried (i.e. `"/busline/44480"`)

    2. The main function will handle the request by calling methods from the busRequest package (i.e. from the file `busLineRequest.go`)

    3. The methods called from the busRequest package will eventually query the external API, in order to retrive information from it and store it in a struct that was declared in the externalAPIResponse package (i.e. The struct `BusLineAPIResponse` in the file `busLineAPIResponse.go`)

    4. The busRequest methods will then instantiate and respond to the client with structs from the busTimingService package, after choosing, filtering, and formatting the appropriate fields from the externalAPIResponse object (i.e. The struct `BusLineWithBuses` from the file `busTimingServiceResponseStructs.go` will be returned)

