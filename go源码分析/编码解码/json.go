package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://localhost:8090/coach?id=1"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	fmt.Println(err, m)

}
