package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Account struct {
	ID       string
	PassWord string
}

//応用

//get

func main() {
	base, _ := url.Parse("https://example.com/")

	reference, _ := url.Parse("index/lists?id=1")

	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)

	req.Header.Add("Content-Type", `application/json"`)

	q := req.URL.Query()

	q.Add("name", "test")

	fmt.Println(q)

	fmt.Println(q.Encode())

	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}

	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

}
