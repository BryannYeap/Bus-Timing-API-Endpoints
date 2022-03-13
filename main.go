package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "encoding/json"
  "github.com/BryannYeap/take_home_assignment/externalAPIResponse"
)


func testArticles(w http.ResponseWriter, r *http.Request) {
  articles := Articles{
    Article{Title: "Test title"},
  }

  json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>HOMEPAGE<h1>")
}

func performGetRequest(url string) {
  response, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  defer response.Body.Close()

  //var responseString strings.Builder
  content, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatal(err)
  }
  //responseString.Write(content)

  var responseObject BusStopAPIResponse
  jsonErr := json.Unmarshal(content, &responseObject)
  if jsonErr != nil {
    fmt.Println(err)
  }

 // fmt.Println(responseString.String())
  fmt.Printf("%+v\n", responseObject);
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/articles", testArticles)
  log.Fatal(http.ListenAndServe(":4000", nil))
}

func main() {
  //handleRequests()
  performGetRequest("https://dummy.uwave.sg/busstop/378204")
}
