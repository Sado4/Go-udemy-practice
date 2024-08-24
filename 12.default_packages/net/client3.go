package main

import (
	"bytes"
	"encoding/json"
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

//Post

func main() {
	base, _ := url.Parse("https://example.com/")

	reference, _ := url.Parse("index/lists?id=1")

	AccountData := &Account{ID: "ABC-DEF", PassWord: "testtest"}
	data, _ := json.Marshal(AccountData)

	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))

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
