package busRequest

import (
    "fmt"
    "net/http"
)

func HomePageRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html> Navigate here to see API end point documentation: " +
        "<a href='https://github.com/BryannYeap/take_home_assignment'> Bus Timing API Endpoints </a></html>")
}