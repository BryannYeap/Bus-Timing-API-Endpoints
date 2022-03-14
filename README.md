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

Response: Provides a link that navigates client to the API Endpoint Documentaiton

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

Response: Gets the id (given) of the bus, and all bus lines that this bus is currently serving. Additionally, for every bus line that this bus is serving, the forecasts for all bus stops on the bus line will be retrieved.
