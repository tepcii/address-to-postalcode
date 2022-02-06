package addrtopos

import (
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
	"net/url"
    "encoding/json"
)

func To(addr string) string {
	req, err := http.NewRequest(http.MethodGet, "https://zipcoda.net/api?address=" + url.QueryEscape(addr), nil)
    if err != nil {
        log.Fatal(err)
    }

	resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }

	body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))

    var stcData Response
    if err := json.Unmarshal([]byte(body), &stcData);
    err != nil {
		log.Fatal(err)
	}
    return stcData.items[0].zipcode
}

type Response struct {
	status int
	length int
    items []ResponseItem
}

type ResponseItem struct {
    zipcode string
    pref string
    components []string
    address string
}