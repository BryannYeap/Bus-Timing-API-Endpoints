package busRequest

import (
    "fmt"
    "os"
    "log"
    "strconv"
    "io/ioutil"
    "net/http"
    "encoding/json"
)

var busLineIDs = []string{"44478", "44479", "44480", "44481"}
var busStopIDs = []string{
    "378204", "383050", "378202", "383049", "382998", "378237", "378233", "378230",
    "378229", "378228", "378227", "382995", "378224", "378226", "383010", "383009",
    "383006", "383004", "378234", "383003", "378222", "383048", "378203", "382999",
    "378225", "383014", "383013", "383011", "377906", "383018", "383015", "378207",
}

func performGetRequest(url string) []byte {
    response, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer response.Body.Close()

    content, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    return content
}

func getEncoder(w http.ResponseWriter) *json.Encoder  {
    encoder := json.NewEncoder(w)
    encoder.SetEscapeHTML(false)
    encoder.SetIndent("", "    ")
    return encoder
}

func convertStringToInt(stringToConvert string) int {
    intToReturn, err := strconv.Atoi(stringToConvert)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return intToReturn
}