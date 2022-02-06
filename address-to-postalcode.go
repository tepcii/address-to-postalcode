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

    if len(stcData.Items) > 0 {
        return stcData.Items[0].Zipcode
    } else {
        return "0000000"
    }
}

type Response struct {
	Status int
	Length int
    Items []ResponseItem
}

type ResponseItem struct {
    Zipcode string
    Pref string
    Components []string
    Address string
}